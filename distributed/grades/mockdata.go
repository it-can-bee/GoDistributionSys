package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "Xucan",
			LastName:  "Zhang",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 90,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 82,
				},
			},
		},
		{
			ID:        2,
			FirstName: "Hailong",
			LastName:  "Li",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 60,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 89,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 82,
				},
			},
		},
	}
}
