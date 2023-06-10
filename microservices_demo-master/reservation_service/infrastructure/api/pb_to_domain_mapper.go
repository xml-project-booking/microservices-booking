package api

/*func ConvertReservationPBToDomain(pbReservationDTO *pb.ReservationDTO) (*domain.ReservationDTO, error) {
	objectID, err := primitive.ObjectIDFromHex(pbReservationDTO.Id)
	if err != nil {
		// Handle the error if the provided ID is not a valid ObjectID
		return nil, err
	}
	accommodationID, err := primitive.ObjectIDFromHex(pbReservationDTO.AccommodationID)
	if err != nil {
		// Handle the error if the provided ID is not a valid ObjectID
		return nil, err
	}
	num, err := strconv.Atoi(pbReservationDTO.GuestNumber)
	// Perform the field mapping
	reservation := &domain.ReservationDTO{
		Id:              objectID,
		GuestNumber:     pbReservationDTO.GuestNumber,
		StartDate:       pbReservationDTO.StartDate,
		EndDate:         pbReservationDTO.EndDate,
		AccommodationID: accommodationID,
		Confirmation:    pbReservationDTO.IsConfirmed,
		// Map other fields as needed
	}

	return reservation, nil
}*/
