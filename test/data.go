package test

import (
	"glintecoTask/entity"
	"time"
)

const AdminUsername = "admin"
const StaffUsername = "staff"
const AdminUuid = "386069f6-72e1-4300-b7a4-a212e728ba5a"
const StaffUuid = "eddc9b4a-d9b5-4189-b291-93874d218805"
const TestValidDeleteUuid = "eddc9b4a-d9b5-4189-b291-93874d218805"
const TestInvalidDeleteUuid = "invalid-uuid"

type MockData struct {
	Admin          entity.User
	Staff          entity.User
	AdminContracts []entity.Contract
	StaffContracts []entity.Contract
}

func (m MockData) AdminContractsUuid() []string {
	return []string{
		"4fb66a7c-f23a-4f4f-9b4e-528c77863df0",
		"7e431d60-e00c-43d2-af49-6ffaf807f4cc",
	}
}

func (m MockData) StaffContractsUuid() []string {
	return []string{
		"0186a127-28fc-4226-9d58-677f6659c9a1",
		"6b62cb16-c317-4181-b386-7a6cbc0e8295",
		"addecce2-bd33-4f7e-9c14-b02abfbdb447",
	}
}

func NewMockData() MockData {
	var m MockData
	createdAt, _ := time.Parse(time.RFC3339, "2024-06-22T16:27:56.946+07:00")
	updatedAt, _ := time.Parse(time.RFC3339, "2024-06-22T16:27:56.946+07:00")
	m.Admin = entity.User{
		Uuid:      AdminUuid,
		Username:  AdminUsername,
		Password:  "12345",
		IsAdmin:   true,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: false,
	}

	createdAt, _ = time.Parse(time.RFC3339, "2024-06-22T16:24:58.743+07:00")
	updatedAt, _ = time.Parse(time.RFC3339, "2024-06-22T16:24:58.743+07:00")
	m.Staff = entity.User{
		Uuid:      StaffUuid,
		Username:  StaffUsername,
		Password:  "12345",
		IsAdmin:   false,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: false,
	}

	acUuid := m.AdminContractsUuid()
	scUuid := m.StaffContractsUuid()
	m.AdminContracts = []entity.Contract{
		{
			Uuid:        acUuid[0],
			Name:        "Admin Contract 1",
			UserUuid:    AdminUuid,
			Description: "Description Admin Contract 1",
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			DeletedAt:   false,
		},
		{
			Uuid:        acUuid[1],
			Name:        "Admin Contract 2",
			UserUuid:    AdminUuid,
			Description: "Description Admin Contract 2",
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			DeletedAt:   false,
		},
	}

	m.StaffContracts = []entity.Contract{
		{
			Uuid:        scUuid[0],
			Name:        "Staff Contract 1",
			UserUuid:    StaffUuid,
			Description: "Description Staff Contract 1",
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			DeletedAt:   false,
		},
		{
			Uuid:        scUuid[1],
			Name:        "Staff Contract 2",
			UserUuid:    StaffUuid,
			Description: "Description Staff Contract 2",
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			DeletedAt:   false,
		},
		{
			Uuid:        scUuid[2],
			Name:        "Staff Contract 3",
			UserUuid:    StaffUuid,
			Description: "Description Staff Contract 3",
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			DeletedAt:   false,
		},
	}

	return m
}
