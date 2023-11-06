-- +goose Up
-- +goose StatementBegin
CREATE TABLE department
ADD COLUMN department_id 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
