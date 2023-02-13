package gorm_test

import (
	"api/src/inventory/domain"
	"api/src/utils"
)

func (s *InventoryIntegrationTestSuite) TestUpdateInventory() {
	req := domain.WriteInventoryBody{Quantity: 10}

	s.Run("Record not found", func() {
		inventory, err := s.Repo.Update(utils.GetDataTypeAddress(100), &req)

		s.Assertions.Nil(inventory)
		s.Assertions.EqualError(err, "not found")
	})

	s.Run("Happy case", func() {
		inventory, err := s.Repo.Update(utils.GetDataTypeAddress(1), &req)

		s.Assertions.NoError(err)
		s.Assertions.Equal(uint(1), inventory.ID)

		updatedInventory := &domain.Inventory{ID: uint(1)}
		result := s.Repo.Conn.First(updatedInventory)
		s.Assertions.NoError(result.Error)
		s.Assertions.Equal(inventory, updatedInventory)
	})
}
