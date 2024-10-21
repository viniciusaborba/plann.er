package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	planner "plann.er/internal/api/spec"
	"plann.er/internal/pgstore"
)

type store interface {
	GetParticipant(ctx context.Context, participantID uuid.UUID) (pgstore.Participant, error)
	ConfirmParticipant(ctx context.Context, participantID uuid.UUID) error
}

type API struct {
	store store
	logger *zap.Logger
	validator *validator.Validate
}

func NewApi(pool *pgxpool.Pool, logger *zap.Logger) API {
	validator := validator.New(validator.WithRequiredStructEnabled())
	return API{pgstore.New(pool), logger, validator}
}

// Confirms a participant on a trip.
// (PATCH /participants/{participantId}/confirm)
func (api *API) PatchParticipantsParticipantIDConfirm(w http.ResponseWriter, r *http.Request, participantID string) *planner.Response {
	id, err := uuid.Parse(participantID)
	if err != nil {
		return planner.PatchParticipantsParticipantIDConfirmJSON400Response(planner.Error{Message: "invalid uuid"})
	}

	participant, err := api.store.GetParticipant(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return planner.PatchParticipantsParticipantIDConfirmJSON400Response(planner.Error{Message: "participant not found"})
		}
		api.logger.Error("failed to get participant", zap.Error(err), zap.String("participant_id", participantID))
		return planner.PatchParticipantsParticipantIDConfirmJSON400Response(planner.Error{Message: "something went wrong"})
	}

	if participant.IsConfirmed {
		return planner.PatchParticipantsParticipantIDConfirmJSON400Response(planner.Error{Message: "participant is already confirmed"})
	}

	if err := api.store.ConfirmParticipant(r.Context(), id); err != nil {
		api.logger.Error("failed to confirm participant", zap.Error(err), zap.String("participant_id", participantID))
		return planner.PatchParticipantsParticipantIDConfirmJSON400Response(planner.Error{Message: "something went wrong"})
	}

	return planner.PatchParticipantsParticipantIDConfirmJSON204Response(nil)
}

// Create a new trip
// (POST /trips)
func (api *API) PostTrips(w http.ResponseWriter, r *http.Request) *planner.Response {
	var body planner.CreateTripRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return planner.PostTripsJSON400Response(planner.Error{Message: "invalid JSON"})
	}
	
	if err := api.validator.Struct(body); err != nil {
		return planner.PostTripsJSON400Response(planner.Error{Message: "invalid input"})
	}

}
// Get a trip details.
// (GET /trips/{tripId})
func (API) GetTripsTripID(w http.ResponseWriter, r *http.Request, tripID string) *planner.Response {
	panic("not implemented") // TODO: Implement
}
// Update a trip.
// (PUT /trips/{tripId})
func (API) PutTripsTripID(w http.ResponseWriter, r *http.Request, tripID string) *planner.Response {
	panic("not implemented") // TODO: Implement
}
// Get a trip activities.
// (GET /trips/{tripId}/activities)
func (API) GetTripsTripIDActivities(w http.ResponseWriter, r *http.Request, tripID string) *planner.Response {
	panic("not implemented") // TODO: Implement
}
// Create a trip activity.
// (POST /trips/{tripId}/activities)
func (API) PostTripsTripIDActivities(w http.ResponseWriter, r *http.Request, tripID string) *planner.Response {
	panic("not implemented") // TODO: Implement
}
// Confirm a trip and send e-mail invitations.
// (GET /trips/{tripId}/confirm)
func (API) GetTripsTripIDConfirm(w http.ResponseWriter, r *http.Request, tripID string) *planner.Response {
	panic("not implemented") // TODO: Implement
}
// Invite someone to the trip.
// (POST /trips/{tripId}/invites)
func (API) PostTripsTripIDInvites(w http.ResponseWriter, r *http.Request, tripID string) *planner.Response {
	panic("not implemented") // TODO: Implement
}
// Get a trip links.
// (GET /trips/{tripId}/links)
func (API) GetTripsTripIDLinks(w http.ResponseWriter, r *http.Request, tripID string) *planner.Response {
	panic("not implemented") // TODO: Implement
}
// Create a trip link.
// (POST /trips/{tripId}/links)
func (API) PostTripsTripIDLinks(w http.ResponseWriter, r *http.Request, tripID string) *planner.Response {
	panic("not implemented") // TODO: Implement
}
// Get a trip participants.
// (GET /trips/{tripId}/participants)
func (API) GetTripsTripIDParticipants(w http.ResponseWriter, r *http.Request, tripID string) *planner.Response {
	panic("not implemented") // TODO: Implement
}
