package model

import (
	"database/sql"
	"time"
)

// ReadAuthListで使用する構造体
type ReadAuthListRequest struct {
	UserID int `json:"user_id"`
}

type ReadAuthListResponse struct {
	ClassName    string `json:"class_name"`
	UserName     string `json:"user_name"`
	DivisionName string `json:"division_name"`
	DocumentID   int    `json:"document_id"`
}

type TakeClassID struct {
	PostID  int    `json:"post_id"`
	ClassID string `json:"class_id"`
}

// ReadDocumentで使用する構造体
type ReadDocumentRequest struct {
	DocumentID int `json:"document_id"`
}

type ReadDocumentResponse struct {
	DocumentID     int       `json:"document_id"`
	RequestAt      time.Time `json:"request_at"`
	StartDate      time.Time `json:"start_date"`
	StartFlame     int       `json:"start_flame"`
	EndDate        time.Time `json:"end_date"`
	EndFlame       int       `json:"end_flame"`
	Location       string    `json:"location"`
	StudentComment string    `json:"student_comment"`
	TeacherComment string    `json:"teacher_comment"`
	UserUuid       string    `json:"user_uuid"`
	ClassName      string    `json:"class_name"`
	UserName       string    `json:"user_name"`
}

// UpdateAuthで使用する構造体
type UpdateAuthRequest struct {
	DocumentID     int    `json:"document_id"`     //ドキュメントID
	UserID         int    `json:"user_id"`         //ユーザID
	TeacherComment string `json:"teacher_comment"` //教員コメント
}

type UpdateDocument struct {
	Status         int    `json:"status"`          // ステータス
	TeacherComment string `json:"teacher_comment"` // 教員コメント
}

// RejectAuthで使用する構造体
type DocumentRejection struct {
	DocumentID int `json:"document_id"`
}

// ReadPositionで使用する構造体
type UserData struct {
	UserID int `json:"user_id"`
}

type UserPosition struct {
	PostID int `json:"post_id"`
}

// CreateDocumentで使用する構造体
type CreateDocumentRequest struct {
	DocumentID        int            `json:"document_id"`         // ドキュメントID
	StudentID         int            `json:"student_id"`          // 学籍番号
	CompanyID         int            `json:"company_id"`          // 会社ID
	ReasonID          int            `json:"reason_id"`           // 理由ID
	RequestDate       time.Time      `json:"request_date"`        // 申請日
	AbsenceStartDate  time.Time      `json:"absence_start_date"`  // 欠席開始日
	AbsenceStartFlame int            `json:"absence_start_flame"` // 欠席開始時限
	AbsenceEndDate    time.Time      `json:"absence_end_date"`    // 欠席終了日
	AbsenceEndFlame   int            `json:"absence_end_flame"`   // 欠席終了時限
	Location          string         `json:"location"`            // 場所
	ReadFlag          bool           `json:"read_flag"`           // 既読フラグ
	Status            int            `json:"status"`              // ステータス
	StudentComment    sql.NullString `json:"student_comment"`     // 学生コメント
	TeacherComment    sql.NullString `json:"teacher_comment"`     // 教員コメント
}

// DeleteDocumentで使用する構造体
type DeleteDocumentRequest struct {
	DocumentID int `json:"document_id"`
}
