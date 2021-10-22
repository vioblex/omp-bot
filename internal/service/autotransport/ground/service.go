package ground

import (
	"errors"
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/autotransport"
)

type DummyGroundService struct {
	list []autotransport.Ground
}

func (s *DummyGroundService) InitEntities() {
	s.list = []autotransport.Ground{
		{Name: "auto", WheelsCount: 4, Color: "black", MaxSpeed: 200},
		{Name: "bus", WheelsCount: 4, Color: "blue", MaxSpeed: 120},
		{Name: "bike", WheelsCount: 2, Color: "red", MaxSpeed: 15},
		{Name: "motorbike", WheelsCount: 2, Color: "black", MaxSpeed: 80},
		{Name: "scooter", WheelsCount: 2, Color: "yellow", MaxSpeed: 25},
	}
}

func (s *DummyGroundService) Describe(groundID uint64) (*autotransport.Ground, error) {
	if err := s.checkSliceIndex(groundID); err != nil {
		return nil, err
	}
	return &s.list[groundID], nil
}

func (s *DummyGroundService) List(cursor uint64, limit uint64) ([]autotransport.Ground, error) {
	if err := s.checkSliceIndex(cursor); err != nil {
		return nil, fmt.Errorf("invalid cursor data: %w", err)
	}
	min := min(cursor+limit, uint64(len(s.list)))
	return s.list[cursor:min], nil
}

func (s *DummyGroundService) Create(ground autotransport.Ground) (uint64, error) {

	err := ground.ValidateFields()
	if err != nil {
		return 0, fmt.Errorf("Error: %s", err)
	}

	s.list = append(s.list, ground)
	return s.Count() - 1, nil
}

func (s *DummyGroundService) Update(groundID uint64, ground autotransport.Ground) error {
	if err := s.checkSliceIndex(groundID); err != nil {
		return err
	}

	origGround := s.list[groundID]
	success := origGround.Copy(ground)

	if success {
		s.list[groundID] = origGround
		return nil
	}

	return fmt.Errorf("invalid fields")
}

func (s *DummyGroundService) Remove(groundID uint64) (bool, error) {
	if err := s.checkSliceIndex(groundID); err != nil {
		return false, err
	}

	s.list = append(s.list[:groundID], s.list[groundID+1:]...)
	return true, nil
}

func (s *DummyGroundService) Count() uint64 {
	return uint64(len(s.list))
}

func (s *DummyGroundService) checkSliceIndex(groundID uint64) error {
	length := s.Count()
	if groundID >= length {
		return errors.New(
			fmt.Sprintf("index %d is out of range 0 - %d", groundID, length),
		)
	}
	return nil
}

func min(a uint64, b uint64) uint64 {
	if a > b {
		return b
	}
	return a
}
