package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/samuskitchen/beer-api-clean-arch/domain"
	"github.com/samuskitchen/beer-api-clean-arch/infrastructure/database"
)

type postgresBeerRepository struct {
	Conn *database.Data
}

// NewBeerRepository constructor
func NewBeerRepository(Connection *database.Data) domain.BeerRepository {
	return &postgresBeerRepository{
		Conn: Connection,
	}
}

func (p *postgresBeerRepository) GetAllBeers(ctx context.Context) ([]domain.Beer, error) {
	rows, err := p.Conn.DB.QueryContext(ctx, selectAllBeers)
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("panic occurred:", err)
		}
	}(rows)

	var beers []domain.Beer
	for rows.Next() {
		var beerRow domain.Beer
		_ = rows.Scan(&beerRow.ID, &beerRow.Name, &beerRow.Brewery, &beerRow.Country, &beerRow.Price, &beerRow.Currency, &beerRow.CreatedAt)
		beers = append(beers, beerRow)
	}

	return beers, nil
}

func (p *postgresBeerRepository) GetBeerById(ctx context.Context, id uint) (domain.Beer, error) {
	row := p.Conn.DB.QueryRowContext(ctx, selectBeerById, id)

	var beerScan domain.Beer
	err := row.Scan(&beerScan.ID, &beerScan.Name, &beerScan.Brewery, &beerScan.Country, &beerScan.Price, &beerScan.Currency, &beerScan.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.Beer{}, nil
		}

		return domain.Beer{}, err
	}

	return beerScan, nil
}

func (p *postgresBeerRepository) CreateBeerWithId(ctx context.Context, beers *domain.Beer) error {
	beer, err := p.GetBeerById(ctx, beers.ID)
	if err != nil {
		return err
	}

	if (domain.Beer{}) != beer {
		return errors.New("beer ID already exists")
	}

	stmt, err := p.Conn.DB.PrepareContext(ctx, insertBeerWithId)
	if err != nil {
		return err
	}

	row := stmt.QueryRowContext(ctx, &beers.ID, &beers.Name, &beers.Brewery, &beers.Country, &beers.Price, &beers.Currency, &beers.CreatedAt)
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println("panic occurred:", err)
		}
	}(stmt)

	err = row.Scan(&beers.ID)
	if err != nil {
		return err
	}

	return nil
}
