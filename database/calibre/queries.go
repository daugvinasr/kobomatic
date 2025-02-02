package calibre

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

// -----------------------------------------------------------------------------

func FindBookCover(db *sql.DB, uuid string) (string, error) {
	query := sq.Select("books.path").
		From("books").
		Where(sq.And{
			sq.NotEq{"books.path": nil},
			sq.Eq{"books.has_cover": true},
			sq.Eq{"books.uuid": uuid},
		})

	row := query.RunWith(db).QueryRow()
	var i string
	err := row.Scan(&i)
	return i, err
}

// -----------------------------------------------------------------------------

type FindBookFilePathRow struct {
	Path string
	Name string
}

func FindBookFilePath(db *sql.DB, uuid string) (FindBookFilePathRow, error) {
	query := sq.Select("books.path", "data.name").
		From("data").
		LeftJoin("books ON data.book = books.id").
		Where(sq.And{
			sq.NotEq{"books.path": nil},
			sq.NotEq{"data.name": nil},
			sq.Eq{"data.format": "EPUB"},
			sq.Eq{"books.uuid": uuid},
		})

	row := query.RunWith(db).QueryRow()
	var i FindBookFilePathRow
	err := row.Scan(&i.Path, &i.Name)
	return i, err
}

// -----------------------------------------------------------------------------

type FindBookMetadataRow struct {
	UUID         string
	Title        string
	Author       string
	Description  string
	Size         int
	Publisher    string
	Series       *string
	SeriesIndex  *float64
	LastModified string
}

func FindBookMetadata(db *sql.DB, uuid string) (FindBookMetadataRow, error) {
	query := sq.
		Select(
			"books.uuid",
			"books.title",
			"authors.name",
			"comments.text",
			"data.uncompressed_size",
			"publishers.name",
			"series.name",
			"books.series_index",
			"books.last_modified",
		).
		From("books").
		LeftJoin("books_authors_link ON books.id = books_authors_link.book").
		LeftJoin("authors ON books_authors_link.author = authors.id").
		LeftJoin("comments ON books.id = comments.book").
		LeftJoin("data ON books.id = data.book").
		LeftJoin("books_publishers_link ON books.id = books_publishers_link.book").
		LeftJoin("publishers ON books_publishers_link.publisher = publishers.id").
		LeftJoin("books_series_link ON books.id = books_series_link.book").
		LeftJoin("series ON books_series_link.series = series.id").
		Where(sq.And{
			sq.NotEq{"books.uuid": nil},
			sq.NotEq{"books.title": nil},
			sq.NotEq{"authors.name": nil},
			sq.NotEq{"comments.text": nil},
			sq.NotEq{"data.uncompressed_size": nil},
			sq.NotEq{"publishers.name": nil},
			sq.NotEq{"books.last_modified": nil},
			sq.Eq{"data.format": "EPUB"},
			sq.Eq{"books.uuid": uuid},
		}).
		GroupBy("books.uuid") // if book has two authors calibre makes two seperate book entries

	row := query.RunWith(db).QueryRow()

	var i FindBookMetadataRow
	err := row.Scan(
		&i.UUID,
		&i.Title,
		&i.Author, &i.Description,
		&i.Size,
		&i.Publisher,
		&i.Series,
		&i.SeriesIndex,
		&i.LastModified,
	)
	return i, err
}

// -----------------------------------------------------------------------------

func FindBooksMetadata(db *sql.DB, lastSync string) ([]FindBookMetadataRow, error) {
	query := sq.
		Select(
			"books.uuid",
			"books.title",
			"authors.name",
			"comments.text",
			"data.uncompressed_size",
			"publishers.name",
			"series.name",
			"books.series_index",
			"books.last_modified",
		).
		From("books").
		LeftJoin("books_authors_link ON books.id = books_authors_link.book").
		LeftJoin("authors ON books_authors_link.author = authors.id").
		LeftJoin("comments ON books.id = comments.book").
		LeftJoin("data ON books.id = data.book").
		LeftJoin("books_publishers_link ON books.id = books_publishers_link.book").
		LeftJoin("publishers ON books_publishers_link.publisher = publishers.id").
		LeftJoin("books_series_link ON books.id = books_series_link.book").
		LeftJoin("series ON books_series_link.series = series.id").
		Where(sq.And{
			sq.NotEq{"books.uuid": nil},
			sq.NotEq{"books.title": nil},
			sq.NotEq{"authors.name": nil},
			sq.NotEq{"comments.text": nil},
			sq.NotEq{"data.uncompressed_size": nil},
			sq.NotEq{"publishers.name": nil},
			sq.NotEq{"books.last_modified": nil},
			sq.Eq{"data.format": "EPUB"},
			sq.Gt{"books.last_modified": lastSync},
		}).
		GroupBy("books.uuid") // if book has two authors calibre makes two seperate book entries

	rows, err := query.RunWith(db).Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindBookMetadataRow
	for rows.Next() {
		var i FindBookMetadataRow
		if err := rows.Scan(
			&i.UUID,
			&i.Title,
			&i.Author,
			&i.Description,
			&i.Size,
			&i.Publisher,
			&i.Series,
			&i.SeriesIndex,
			&i.LastModified,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
