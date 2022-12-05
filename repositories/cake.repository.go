package repositories

import (
	contextmanager "cake_store/context_manager"
	"cake_store/entities"
	"cake_store/handler"
	"context"
	"errors"
	"time"
)

func GetAllCake() ([]entities.Cake, error) {
	db := contextmanager.OpenDBConnection()
	defer db.Close()

	var cakes []entities.Cake
	results, err := db.Query("SELECT * from cake ORDER BY rating DESC, title")

	if err != nil {
		err := errors.New("Error when querying to databases")
		err = handler.Wrap(err, "repositories")
		return nil, err
	}

	for results.Next() {
		var cake entities.Cake
		// for each row, scan the result into our cake composite object
		err = results.Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		if err != nil {
			err := errors.New("Error when processing data result")
			err = handler.Wrap(err, "repositories")
			return nil, err
		}
		cakes = append(cakes, cake)
	}

	return cakes, nil
}

func CreateCake(ctx context.Context, cake entities.Cake) error {
	db := contextmanager.OpenDBConnection()
	defer db.Close()

	transaction, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer transaction.Rollback()

	_, err = transaction.ExecContext(ctx, "INSERT INTO cake (title, description, rating, image, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)", cake.Title, cake.Description, cake.Rating, cake.Image, time.Now(), time.Now())
	if err != nil {
		return err
	}

	// Commit the transaction.
	if err = transaction.Commit(); err != nil {
		return err
	}
	return nil
}

func GetOneCake(cakeId int) (entities.Cake, error) {
	var cake entities.Cake
	db := contextmanager.OpenDBConnection()
	defer db.Close()

	err := db.QueryRow("SELECT * from cake WHERE id=?", cakeId).Scan(&cake.ID, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
	if err != nil {
		return cake, err
	}

	return cake, nil
}

func UpdateCake(ctx context.Context, cakeId int, cake entities.Cake) error {
	db := contextmanager.OpenDBConnection()
	defer db.Close()

	transaction, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer transaction.Rollback()

	_, err = transaction.ExecContext(ctx, "UPDATE cake SET title=?, description=?, rating=?, image=?, updated_at=? WHERE id=?", cake.Title, cake.Description, cake.Rating, cake.Image, time.Now(), cakeId)
	if err != nil {
		return err
	}

	// Commit the transaction.
	if err = transaction.Commit(); err != nil {
		return err
	}
	return nil
}

func DeleteCake(ctx context.Context, cakeId int) error {
	// Open database connection
	db := contextmanager.OpenDBConnection()
	defer db.Close()

	// Start transaction
	transaction, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer transaction.Rollback()

	// Execute Query
	_, err = transaction.ExecContext(ctx, "DELETE FROM cake where id=?", cakeId)
	if err != nil {
		return err
	}

	// Commit the transaction.
	if err = transaction.Commit(); err != nil {
		return err
	}
	return nil
}
