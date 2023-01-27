// Package exchange provides data types and methods that simulate a stock exchange capable of
// handling buy, sell and cancel orders.
package exchange

import (
	"fmt"

	"github.com/angusgmorrison/sedgewick_algorithms/struct/heap"
	"github.com/google/uuid"
)

type OrderType int

const (
	OrderTypeBuy OrderType = iota + 1
	OrderTypeSell
)

type TickerSymbol string

// order describes the methods common to all orders.
type order interface {
	ID() uuid.UUID
	Type() OrderType
	TickerSymbol() TickerSymbol
}

// orderData contains the fields common to all orders.
type orderData struct {
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

func (o *orderData) quantityOutstanding() uint64 {
	return o.quantityOrdered - o.quantityFilled
}

func (o *orderData) isFilled() bool {
	return o.quantityOrdered == o.quantityFilled
}

type buyOrder struct {
	orderData
}

var _ order = (*buyOrder)(nil)
var _ heap.Prioritizable[*buyOrder] = (*buyOrder)(nil)

func (bo *buyOrder) ID() uuid.UUID {
	return bo.id
}

func (bo *buyOrder) Type() OrderType {
	return OrderTypeBuy
}

func (bo *buyOrder) TickerSymbol() TickerSymbol {
	return bo.tickerSymbol
}

// HasPriority prioritizes buyOrders with the highest bid followed by the lowest sequence ID.
func (a *buyOrder) HasPriority(b *buyOrder) bool {
	if a.priceCents == b.priceCents {
		return a.sequenceID < b.sequenceID
	}

	return a.priceCents > b.priceCents
}

type sellOrder struct {
	orderData
}

var _ order = (*sellOrder)(nil)

func (so *sellOrder) ID() uuid.UUID {
	return so.id
}

func (so *sellOrder) Type() OrderType {
	return OrderTypeSell
}

func (so *sellOrder) TickerSymbol() TickerSymbol {
	return so.tickerSymbol
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
	sellOrderHeaps map[TickerSymbol]*heap.SymbolHeap[uuid.UUID, *sellOrder]
	// sellOrderHeaps maps a TickerSymbol to a heap of all unfilled buy orders for that stock.
	buyOrderHeaps map[TickerSymbol]*heap.SymbolHeap[uuid.UUID, *buyOrder]
	// ledger maps an ID to an order for constant-time lookup of any order, regardless of type.
	ledger map[uuid.UUID]order
}

func newBook() *book {
	return &book{
		sellOrderHeaps: make(map[TickerSymbol]*heap.SymbolHeap[uuid.UUID, *sellOrder]),
		buyOrderHeaps:  make(map[TickerSymbol]*heap.SymbolHeap[uuid.UUID, *buyOrder]),
		ledger:         make(map[uuid.UUID]order),
	}
}

type SequenceError struct {
	wantSequenceID uint64
	order          orderData
}

func (e *SequenceError) Error() string {
	return fmt.Sprintf("want Order with SequenceID %d, got %+v", e.wantSequenceID, e.order)
}

func (b *book) executeBuyOrder(bo *buyOrder) error {
	if bo.sequenceID != b.lastSequenceID+1 {
		return &SequenceError{
			wantSequenceID: b.lastSequenceID + 1,
			order:          bo.orderData,
		}
	}
	b.lastSequenceID++

	sellOrdersForStock, ok := b.sellOrderHeaps[bo.tickerSymbol]
	if !ok || sellOrdersForStock.IsEmpty() { // no sell orders exist, book the buy order for later
		b.bookBuyOrder(bo)
		return nil // notify counterparties that order has been registered
	}

	_, so, _ := sellOrdersForStock.Pop()
	for bo.priceCents >= so.priceCents {
		b.fillOrderPair(bo, so)

		if bo.isFilled() {
			if !so.isFilled() {
				sellOrdersForStock.Push(so.id, so) // TODO: Peek and only pop if order is filled
			}
			return nil // notify counterparties...
		}

		if _, so, ok = sellOrdersForStock.Pop(); !ok { // no more sell orders
			break
		}
	}

	// Unable to fill buy order at the currently available prices.
	b.bookBuyOrder(bo)
	return nil // notify...
}

func (b *book) executeSellOrder(so *sellOrder) error {
	if so.sequenceID != b.lastSequenceID+1 {
		return &SequenceError{
			wantSequenceID: b.lastSequenceID + 1,
			order:          so.orderData,
		}
	}
	b.lastSequenceID++

	buyOrdersForStock, ok := b.buyOrderHeaps[so.tickerSymbol]
	if !ok || buyOrdersForStock.IsEmpty() {
		b.bookSellOrder(so)
		return nil // notify...
	}

	_, bo, _ := buyOrdersForStock.Pop() // TODO: Peek and only pop if order is filled
	for so.priceCents <= bo.priceCents {
		b.fillOrderPair(bo, so)

		if so.isFilled() {
			if !bo.isFilled() {
				buyOrdersForStock.Push(bo.id, bo)
			}
			return nil // notify
		}

		if _, bo, ok = buyOrdersForStock.Pop(); !ok { // no more buy orders
			break
		}
	}

	// Unable to fill sell order at the currently available prices.
	b.bookSellOrder(so)
	return nil // notify
}

func (b *book) bookBuyOrder(bo *buyOrder) {
	buyOrdersForStock, ok := b.buyOrderHeaps[bo.tickerSymbol]
	if !ok {
		buyOrdersForStock = heap.NewSymbolHeap[uuid.UUID, *buyOrder]()
		b.buyOrderHeaps[bo.tickerSymbol] = buyOrdersForStock
	}

	b.ledger[bo.id] = bo
	buyOrdersForStock.Push(bo.id, bo)
}

func (b *book) bookSellOrder(so *sellOrder) {
	sellOrdersForStock, ok := b.sellOrderHeaps[so.tickerSymbol]
	if !ok {
		sellOrdersForStock = heap.NewSymbolHeap[uuid.UUID, *sellOrder]()
		b.sellOrderHeaps[so.tickerSymbol] = sellOrdersForStock
	}

	b.ledger[so.id] = so
	sellOrdersForStock.Push(so.id, so)
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
	order, ok := b.ledger[id]
	if !ok {
		return &OrderNotInLedgerError{id: id}
	}

	switch order.Type() {
	case OrderTypeBuy:
		b.cancelBuyOrder(id, order.TickerSymbol())
	case OrderTypeSell:
		b.cancelSellOrder(id, order.TickerSymbol())
	default:
		panic(fmt.Errorf("unknown order type %d", order.Type()))
	}
	return nil
}

type OrderNotFoundError struct {
	orderID uuid.UUID
}

func (e *OrderNotFoundError) Error() string {
	return fmt.Sprintf("order with ID %s not found", e.orderID)
}

func (b *book) cancelSellOrder(id uuid.UUID, sym TickerSymbol) error {
	sellOrdersForStock, ok := b.sellOrderHeaps[sym]
	if !ok || !sellOrdersForStock.Contains(id) {
		return &OrderNotFoundError{orderID: id}
	}

	sellOrdersForStock.Delete(id)
	return nil
}

func (b *book) cancelBuyOrder(id uuid.UUID, sym TickerSymbol) error {
	buyOrdersForStock, ok := b.buyOrderHeaps[sym]
	if !ok || !buyOrdersForStock.Contains(id) {
		return &OrderNotFoundError{orderID: id}
	}

	buyOrdersForStock.Delete(id)
	return nil
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}

	return b
}
