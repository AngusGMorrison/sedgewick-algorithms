// Package exchange provides data types and methods that simulate a stock exchange capable of
// handling buy, sell and cancel orders.
package exchange

import (
	"fmt"

	"github.com/angusgmorrison/sedgewick_algorithms/struct/heap"
	"github.com/google/uuid"
)

type TickerSymbol string

type order struct {
	// Unique identifier.
	id uuid.UUID
	// Monotonic sequence number.
	sequenceID uint64
	// counterpartyID  string  // unused but desirable in real-world scenario
	tickerSymbol    TickerSymbol
	priceCents      uint64
	quantityOrdered uint64
	quantityFilled  uint64
}

func (o *order) quantityOutstanding() uint64 {
	return o.quantityOrdered - o.quantityFilled
}

func (o *order) isFilled() bool {
	return o.quantityOrdered == o.quantityFilled
}

type buyOrder struct {
	order
}

var _ heap.Prioritizable[*buyOrder] = (*buyOrder)(nil)

// HasPriority prioritizes buyOrders with the highest bid followed by the lowest sequence ID.
func (a *buyOrder) HasPriority(b *buyOrder) bool {
	if a.priceCents == b.priceCents {
		return a.sequenceID < b.sequenceID
	}

	return a.priceCents > b.priceCents
}

type sellOrder struct {
	order
}

var _ heap.Prioritizable[*sellOrder] = (*sellOrder)(nil)

// HasPriority prioritizes sellOrders with the lowest ask followed by the lowest sequence ID.
func (a *sellOrder) HasPriority(b *sellOrder) bool {
	if a.priceCents == b.priceCents {
		return a.sequenceID < b.sequenceID
	}

	return a.priceCents < b.priceCents
}

// Exchange represents an object that behaves like a stock exchange.
type Exchange interface {
	Buy(order *buyOrder) error
	Sell(order *sellOrder) error
	Cancel(orderID uuid.UUID) error
}

type exchange struct {
	book *book
	// other fields may include clients, db conns, etc. for persisting orders, notifying
	// counterparties, etc.
}

var _ Exchange = (*exchange)(nil)

func NewExchange() Exchange {
	return &exchange{
		book: newBook(),
	}
}

// TODO: Accept a BuyOrderRequest and transform into a buyOrder, performing sequencing internally.
func (b *exchange) Buy(bo *buyOrder) error {
	// Perform pre-actions (e.g. sequencing).
	if err := b.book.executeBuyOrder(bo); err != nil {
		// Handle error.
		return err
	}
	// Perform post-actions (e.g. notify counterparties if trade is completed).
	return nil
}

// TODO: Accept a SellOrderRequest and transform into a buyOrder, performing sequencing internally.
func (b *exchange) Sell(so *sellOrder) error {
	// Pre-actions.
	if err := b.book.executeSellOrder(so); err != nil {
		// Handle err.
		return err
	}
	// Post-actions.
	return nil
}

func (b *exchange) Cancel(orderID uuid.UUID) error {
	// Pre-actions.
	if err := b.book.cancelOrder(orderID); err != nil {
		// Handle err.
		return err
	}
	// Post-actions.
	return nil
}

type book struct {
	// lastSequenceID is the sequence ID of the last order processed. It is used to prevent lost
	// orders and ensure monotonic order execution.
	lastSequenceID uint64
	// sellOrderHeaps maps a TickerSymbol to a heap of all unfilled sell orders for that stock.
	sellOrderHeaps map[TickerSymbol]*heap.IndexedHeap[*sellOrder]
	// sellOrderHeaps maps a TickerSymbol to a heap of all unfilled buy orders for that stock.
	buyOrderHeaps map[TickerSymbol]*heap.IndexedHeap[*buyOrder]
	// sellOrderLedger maps each sellOrder ID to an indexed order, allowing constant-time find
	// operations in the order's corresponding heap. This supports order cancellation in logarithmic
	// time.
	sellOrderLedger map[uuid.UUID]*heap.IndexedEntry[*sellOrder]
	// buyOrderLedger maps each buyOrder ID to an indexed order.
	buyOrderLedger map[uuid.UUID]*heap.IndexedEntry[*buyOrder]
}

func newBook() *book {
	return &book{
		sellOrderHeaps:  make(map[TickerSymbol]*heap.IndexedHeap[*sellOrder]),
		buyOrderHeaps:   make(map[TickerSymbol]*heap.IndexedHeap[*buyOrder]),
		sellOrderLedger: make(map[uuid.UUID]*heap.IndexedEntry[*sellOrder]),
		buyOrderLedger:  make(map[uuid.UUID]*heap.IndexedEntry[*buyOrder]),
	}
}

type SequenceError struct {
	wantSequenceID uint64
	order          order
}

func (e *SequenceError) Error() string {
	return fmt.Sprintf("want Order with SequenceID %d, got %+v", e.wantSequenceID, e.order)
}

func (b *book) executeBuyOrder(bo *buyOrder) error {
	if bo.sequenceID != b.lastSequenceID+1 {
		return &SequenceError{
			wantSequenceID: b.lastSequenceID + 1,
			order:          bo.order,
		}
	}
	b.lastSequenceID++

	sellOrdersForStock, ok := b.sellOrderHeaps[bo.tickerSymbol]
	if !ok || sellOrdersForStock.IsEmpty() { // no sell orders exist, book the buy order for later
		if err := b.bookBuyOrder(bo); err != nil {
			return err
		}
		return nil // notify counterparties that order has been registered
	}

	so, _ := sellOrdersForStock.Pop()
	for bo.priceCents >= so.priceCents {
		b.fillOrderPair(bo, so)

		if so.isFilled() {
			delete(b.sellOrderLedger, so.id)
		}

		if bo.isFilled() {
			if !so.isFilled() {
				sellOrdersForStock.Push(so)
			}
			return nil // notify counterparties...
		}

		if so, ok = sellOrdersForStock.Pop(); !ok { // no more sell orders
			break
		}
	}

	// Unable to fill buy order at the currently available prices.
	if err := b.bookBuyOrder(bo); err != nil {
		return err
	}
	return nil // notify...
}

func (b *book) executeSellOrder(so *sellOrder) error {
	if so.sequenceID != b.lastSequenceID+1 {
		return &SequenceError{
			wantSequenceID: b.lastSequenceID + 1,
			order:          so.order,
		}
	}
	b.lastSequenceID++

	buyOrdersForStock, ok := b.buyOrderHeaps[so.tickerSymbol]
	if !ok || buyOrdersForStock.IsEmpty() {
		if err := b.bookSellOrder(so); err != nil {
			return err
		}
		return nil // notify...
	}

	bo, _ := buyOrdersForStock.Pop()
	for so.priceCents <= bo.priceCents {
		b.fillOrderPair(bo, so)

		if bo.isFilled() {
			delete(b.buyOrderLedger, bo.id)
		}

		if so.isFilled() {
			if !bo.isFilled() {
				buyOrdersForStock.Push(bo)
			}
			return nil // notify
		}

		if bo, ok = buyOrdersForStock.Pop(); !ok { // no more buy orders
			break
		}
	}

	// Unable to fill sell order at the currently available prices.
	if err := b.bookSellOrder(so); err != nil {
		return err
	}
	return nil // notify
}

// OrderDoubleBookedError indicates an attempt to book the same order twice, which should never
// happen.
type OrderDoubleBookedError struct {
	order order
}

func (e *OrderDoubleBookedError) Error() string {
	return fmt.Sprintf("attempted to book order that was already in ledger: %+v", e.order)
}

func (b *book) bookBuyOrder(bo *buyOrder) error {
	if _, ok := b.buyOrderLedger[bo.id]; ok {
		return &OrderDoubleBookedError{order: bo.order}
	}

	buyOrdersForStock, ok := b.buyOrderHeaps[bo.tickerSymbol]
	if !ok {
		buyOrdersForStock = heap.NewIndexedHeap[*buyOrder]()
		b.buyOrderHeaps[bo.tickerSymbol] = buyOrdersForStock
	}

	indexedBO := buyOrdersForStock.Push(bo)
	b.buyOrderLedger[bo.id] = indexedBO

	return nil
}

func (b *book) bookSellOrder(so *sellOrder) error {
	if _, ok := b.sellOrderLedger[so.id]; ok {
		return &OrderDoubleBookedError{order: so.order}
	}

	sellOrdersForStock, ok := b.sellOrderHeaps[so.tickerSymbol]
	if !ok {
		sellOrdersForStock = heap.NewIndexedHeap[*sellOrder]()
		b.sellOrderHeaps[so.tickerSymbol] = sellOrdersForStock
	}

	indexedSO := sellOrdersForStock.Push(so)
	b.sellOrderLedger[so.id] = indexedSO

	return nil
}

func (b *book) fillOrderPair(bo *buyOrder, so *sellOrder) {
	maxFillable := min(bo.quantityOutstanding(), so.quantityOutstanding())
	bo.quantityFilled += maxFillable
	so.quantityFilled += maxFillable
}

type OrderNotInLedgerError struct {
	id uuid.UUID
}

func (e *OrderNotInLedgerError) Error() string {
	return fmt.Sprintf("no matching order for ID %s in buy or sell ledgers", e.id)
}

func (b *book) cancelOrder(id uuid.UUID) error {
	indexedSO, ok := b.sellOrderLedger[id]
	if ok {
		return b.cancelSellOrderAtIndex(indexedSO.Entry(), indexedSO.Index())
	}

	indexedBO, ok := b.buyOrderLedger[id]
	if ok {
		return b.cancelBuyOrderAtIndex(indexedBO.Entry(), indexedBO.Index())
	}

	return &OrderNotInLedgerError{id: id}
}

type OrderNotInHeapError struct {
	order order
	index int
}

func (e *OrderNotInHeapError) Error() string {
	return fmt.Sprintf("order with ID %s not found at heap index %d: %+v", e.order.id, e.index, e.order)
}

func (b *book) cancelSellOrderAtIndex(so *sellOrder, i int) error {
	sellOrdersForStock, ok := b.sellOrderHeaps[so.tickerSymbol]
	if !ok {
		return &OrderNotInHeapError{order: so.order, index: i}
	}

	if _, ok := sellOrdersForStock.Remove(i); !ok {
		return &OrderNotInHeapError{order: so.order, index: i}
	}

	return nil
}

func (b *book) cancelBuyOrderAtIndex(bo *buyOrder, i int) error {
	buyOrdersForStock, ok := b.buyOrderHeaps[bo.tickerSymbol]
	if !ok {
		return &OrderNotInHeapError{order: bo.order, index: i}
	}

	if _, ok := buyOrdersForStock.Remove(i); !ok {
		return &OrderNotInHeapError{order: bo.order, index: i}
	}

	return nil
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}

	return b
}
