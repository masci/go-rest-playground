package main

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/masci/go-rest-playground/models"
)

// ListClasses handles GET requests at /classes
func ListClasses(w http.ResponseWriter, r *http.Request) {
	list := []render.Renderer{}
	classes, err := storage.GetClasses()
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	// Get all the classes from the storage and render them one after the other
	// using the RenderList helper from the Chi framework
	for _, c := range classes {
		list = append(list, NewClassResponse(c))
	}

	if err := render.RenderList(w, r, list); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

// CreateClass handles POST requests at /classes
func CreateClass(w http.ResponseWriter, r *http.Request) {
	data := &ClassPayload{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	c := data.Class
	storage.AddClass(c)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewClassResponse(c))
}

// GetClass handles GET requests at /classes/<CLASS_ID>
func GetClass(w http.ResponseWriter, r *http.Request) {
	// get the Class object from the request context
	class := r.Context().Value("class").(*models.Class)

	if err := render.Render(w, r, NewClassResponse(class)); err != nil {
		render.Render(w, r, ErrRender(err))
	}
}

// UpdateClass handles PUT requests at /classes/<CLASS_ID>
func UpdateClass(w http.ResponseWriter, r *http.Request) {
	// get the Class object from the request context
	class := r.Context().Value("class").(*models.Class)

	// render the payload to see if there's all we need to build a Class object.
	// In a real-world scenario we could also enrich the object with some metadata
	// if needed.
	data := &ClassPayload{Class: class}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	class = data.Class

	// persist the changes
	storage.UpdateClass(class.ID, class)

	// render the updated Class
	render.Render(w, r, NewClassResponse(class))
}

// DeleteClass handles DELETE requests at /classes/<CLASS_ID>
func DeleteClass(w http.ResponseWriter, r *http.Request) {
	// get the Class object from the request context
	class := r.Context().Value("class").(*models.Class)

	if err := storage.DeleteClass(class.ID); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, NewClassResponse(class))
}

// ListBookings handles GET requests at /bookings
func ListBookings(w http.ResponseWriter, r *http.Request) {
	list := []render.Renderer{}
	bookings, err := storage.GetBookings()
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	// Get all the bookings from the storage and render them one after the other
	// using the RenderList helper from the Chi framework
	for _, b := range bookings {
		list = append(list, NewBookingResponse(b))
	}

	if err := render.RenderList(w, r, list); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

// CreateBooking handles POST requests at /bookings
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	data := &BookingPayload{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	// persist booking
	b := data.Booking
	if _, err := storage.AddBooking(b); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, NewBookingResponse(b))
}

// GetBooking handles GET requests at /bookings/<BOOKING_ID>
func GetBooking(w http.ResponseWriter, r *http.Request) {
	// get the Booking object from the request context
	booking := r.Context().Value("booking").(*models.Booking)

	if err := render.Render(w, r, NewBookingResponse(booking)); err != nil {
		render.Render(w, r, ErrRender(err))
	}
}

// UpdateBooking handles PUT requests at /bookings/<BOOKING_ID>
func UpdateBooking(w http.ResponseWriter, r *http.Request) {
	// get the Booking object from the request context
	booking := r.Context().Value("booking").(*models.Booking)

	// render the payload to see if there's all we need to build a Booking object.
	// In a real-world scenario we could also enrich the object with some metadata
	// if needed.
	data := &BookingPayload{Booking: booking}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	booking = data.Booking

	// persist the changes
	storage.UpdateBooking(booking.ID, booking)

	// render the updated Booking
	render.Render(w, r, NewBookingResponse(booking))
}

// DeleteBooking handles DELETE requests at /bookings/<BOOKING_ID>
func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	// get the Class object from the request context
	booking := r.Context().Value("class").(*models.Booking)

	if err := storage.DeleteBooking(booking.ID); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, NewBookingResponse(booking))
}

/*
	Request/Response types.

	For the sake of simplicity, we're using one struct to represent both
	Requests and Responses for a single resource (Class and Booking) but
	if needed these types can be further specialized.

	Please note within the scope of this exercise Render and Bind methods
	don't really need to do anything but we're ready to support more advanced
	usage.
*/

// ClassPayload represents Request and Response payload for the Class resource
type ClassPayload struct {
	*models.Class
}

// NewClassResponse returns a ClassPayload object
func NewClassResponse(class *models.Class) *ClassPayload {
	return &ClassPayload{class}
}

// Render is a no-op for our use case
func (cp *ClassPayload) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Bind only ensures the Class object can be created in our use case
func (cp *ClassPayload) Bind(r *http.Request) error {
	// cp.Class is nil when there is no field in the request
	if cp.Class == nil {
		return errors.New("missing required Class object")
	}

	return nil
}

// BookingPayload represents Request and Response payload for the Class resource
type BookingPayload struct {
	*models.Booking
}

// NewBookingResponse returns a ClassPayload object
func NewBookingResponse(booking *models.Booking) *BookingPayload {
	return &BookingPayload{booking}
}

// Render is a no-op for our use case
func (cp *BookingPayload) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Bind only ensures the Class object can be created in our use case
func (cp *BookingPayload) Bind(r *http.Request) error {
	// cp.Booking is nil when there is no field in the request
	if cp.Booking == nil {
		return errors.New("missing required Booking object")
	}

	return nil
}
