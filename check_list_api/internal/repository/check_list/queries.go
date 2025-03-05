package check_list

const (
	queryCreateTask = `
		INSERT INTO task.tasks (
			user_id,
			title,
			description,
			done
		)
		VALUES ($1, $2, $3, $4)
		RETURNING task_id
	`

	queryGetTaskByUserIDAndTaskID = `
		SELECT task_id,
		    title,
		    description,
			done
		FROM task.tasks
		WHERE user_id = $1
			AND task_id = $2
	`

	queryGetAllTasksByUserID = `
		SELECT task_id,
		    title,
		    description,
			done
		FROM task.tasks
		WHERE user_id = $1
	`

	queryDeleteTask = `
		DELETE FROM task.tasks
		WHERE task_id = $1
			AND user_id = $2
		RETURNING task_id
	`

	queryDoneTask = `
		UPDATE task.tasks
		SET done = $1
		WHERE task_id = $2
		  	AND user_id = $3
		RETURNING task_id
	`
)
