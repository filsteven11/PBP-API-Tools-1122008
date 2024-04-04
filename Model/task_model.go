package model

type Task struct {
	ID        int
	Name      string
	StudentID int
}

func NewTask(id int, name string, studentID int) *Task {
	return &Task{
		ID:        id,
		Name:      name,
		StudentID: studentID,
	}
}
