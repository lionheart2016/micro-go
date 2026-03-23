package mock

import (
	"micro-go/internal/model"
)

// UserMock 用户 Mock 数据
type UserMock struct {
	users map[int]*model.User
	nextID int
}

// NewUserMock 创建新的用户 Mock
func NewUserMock() *UserMock {
	return &UserMock{
		users:  make(map[int]*model.User),
		nextID: 1,
	}
}

// Create 创建用户
func (m *UserMock) Create(user *model.User) error {
	user.ID = m.nextID
	m.users[user.ID] = user
	m.nextID++
	return nil
}

// GetByID 根据 ID 获取用户
func (m *UserMock) GetByID(id int) (*model.User, error) {
	user, exists := m.users[id]
	if !exists {
		return nil, nil
	}
	return user, nil
}

// GetAll 获取所有用户
func (m *UserMock) GetAll() ([]*model.User, error) {
	users := make([]*model.User, 0, len(m.users))
	for _, user := range m.users {
		users = append(users, user)
	}
	return users, nil
}

// Update 更新用户
func (m *UserMock) Update(id int, user *model.User) error {
	_, exists := m.users[id]
	if !exists {
		return nil
	}
	user.ID = id
	m.users[id] = user
	return nil
}

// Delete 删除用户
func (m *UserMock) Delete(id int) error {
	delete(m.users, id)
	return nil
}
