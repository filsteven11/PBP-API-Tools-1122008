package Controller

import (
	"fmt"
)

type TaskController struct {
	EmailCtrl *EmailController
	CacheCtrl *CacheController
}

func (tc *TaskController) SubmitTaskNotification(email, taskName string) error {
	// Menyiapkan pesan email
	subject := "Notifikasi Tugas"
	body := fmt.Sprintf("Anda memiliki tugas baru: %s", taskName)

	// Mengirim email notifikasi
	err := tc.EmailCtrl.SendEmail(email, subject, body)
	if err != nil {
		return err
	}

	return nil
}
