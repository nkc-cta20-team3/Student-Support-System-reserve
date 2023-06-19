package model

import (
	"database/sql"
	"time"
)

type AbsenceDocument struct {
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

// UnAuthorization 引数の構造体
type TeacherData struct {
	TeacherID int `json:"teacher_id"` //教員ID
}

// UnAuthorization 戻り値の構造体
type UnAuthorizeList struct {
	ClassName       string `json:"class_name"`       //クラス名称
	StudentName     string `json:"student_name"`     //学生氏名
	AbsenceCategory string `json:"absence_category"` //種別
	DocumentID      int    `json:"document_id"`      //書類ID
}

// UnAuthorization,UnAuthorizationGET 作業用構造体
type TakeClassName struct {
	PositionID int    `json:"position_id"`
	ClassName  string `json:"class_name"`
}
