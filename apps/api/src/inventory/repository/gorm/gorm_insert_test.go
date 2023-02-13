package gorm_test

import "api/src/inventory/domain"

func (s *InventoryIntegrationTestSuite) TestInsertInventory() {
	s.Run("Happy case", func() {
		req := domain.WriteInventoryBody{Quantity: 5}
		inventory, err := s.Repo.Insert(&req)

		s.Assertions.NoError(err)
		s.Assertions.NotNil(inventory)
		s.Assertions.Equal(req.Quantity, inventory.Quantity)
	})
}
