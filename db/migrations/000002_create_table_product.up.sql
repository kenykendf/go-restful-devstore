CREATE TABLE IF NOT EXISTS product (
    ID bigserial not null primary key,
    name varchar(50),
    description text,
    currency varchar(50),
    price bigint,
    total_stock integer not null,
    is_active boolean default false,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    category_id bigint not null,
    constraint fk_category foreign key (category_id) references categories(id) 
);