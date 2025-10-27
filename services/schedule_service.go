package services

import (
	"consultancy_hours/models"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var allDailySlots []string

func init() {
	allDailySlots = make([]string, 24)
	for i := 0; i < 24; i++ {
		allDailySlots[i] = fmt.Sprintf("%02d:00", i)
	}
}

type ScheduleService struct {
	collection *mongo.Collection
}

func NewScheduleService(db *mongo.Database) *ScheduleService {
	return &ScheduleService{
		collection: db.Collection("scheduling"),
	}
}

func (s *ScheduleService) getBookedSchedules() (map[string]bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	bookedSlots := make(map[string]bool)

	for cursor.Next(ctx) {
		var appointment models.Schedule
		if err := cursor.Decode(&appointment); err != nil {
			return nil, err
		}
		bookedSlots[appointment.IDTime] = true
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return bookedSlots, nil
}

func (s *ScheduleService) GetAvailableSchedules() ([]models.AvailableSlot, error) {

	bookedSlots, err := s.getBookedSchedules()
	if err != nil {
		return nil, err
	}

	availableSlots := make([]models.AvailableSlot, 0)

	for _, slotID := range allDailySlots {
		if _, isBooked := bookedSlots[slotID]; !isBooked {
			availableSlots = append(availableSlots, models.AvailableSlot{
				IDTime: slotID,
				Status: "disponivel",
			})
		}
	}

	return availableSlots, nil
}

func (s *ScheduleService) CreateSchedule(schedule models.Schedule) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	isValidSlot := false
	for _, slot := range allDailySlots {
		if slot == schedule.IDTime {
			isValidSlot = true
			break
		}
	}
	if !isValidSlot {
		return nil, errors.New("id_horario invalid")
	}

	var existing models.Schedule
	err := s.collection.FindOne(ctx, bson.M{"id_horario": schedule.IDTime}).Decode(&existing)

	if err == nil {
		return nil, errors.New("schedule unavailable")
	}
	if err != mongo.ErrNoDocuments {
		return nil, err
	}

	res, err := s.collection.InsertOne(ctx, schedule)
	if err != nil {
		return nil, err
	}
	return res, nil
}
