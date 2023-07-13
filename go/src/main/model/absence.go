package model

import (
	"database/sql"
	"time"
)

// ReadAuthListで使用する構造体
type ReadAuthListRequest struct {
	TeacherID int `json:"teacher_id"`
}

type ReadAuthListResponse struct {
	ClassName  string `json:"class_name"`
	StudentName string `json:"student_name"`
	AbsenceCategory string `json:"absence_category"`
	DocumentID int `json:"document_id"`
}

type TakeClassName struct {
	PositionID int    `json:"position_id"`
	ClassName  string `json:"class_name"`
}



// ReadDocumentで使用する構造体
type ReadDocumentRequest struct {
	DocumentID int `json:"document_id"`
}

type ReadDocumentResponse struct {
	DocumentID int `json:"document_id"`
	RequestDate time.Time `json:"request_date"`
	StudentID int `json:"student_id"`
	ClassName string `json:"class_name"`
	StudentName string `json:"student_name"`
	StartDate time.Time `json:"absence_start_date"`
	StartFlame int `json:"start_flame"`
	EndDate time.Time `json:"end_date"`
	EndFlame int `json:"end_flame"`
	Location string `json:"location"`
	StudentComment string `json:"student_comment"`
	TeacherComment string `json:"teacher_comment"`
}

// UpdateAuthで使用する構造体
type UpdateAuthRequest struct {
	DocumentID int `json:"document_id"`
	TeacherID int `json:"teacher_id"`
	TeacherComment string `json:"teacher_comment"`
}

type UpdateDocument struct {
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
	TeacherComment    string		 `json:"teacher_comment"`     // 教員コメント
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

// Rejectionで使用する構造体
type DocumentRejection struct{
	DocumentID int `json:"document_id"`
}