package reactive

// UPDATED - Event emitted when the value of a cell is updated
const UPDATED = "UPDATED"

// Reactive - TODO
type Reactive interface {
	emit(event string)
	on(event string, callback func())
}
