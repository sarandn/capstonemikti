CREATE TABLE Ticket (
    ticket_id serial PRIMARY KEY,
    event_id_fk INT,
    ticket_type VARCHAR(255),
    ticket_price DECIMAL(10, 2),
    quantity_avail INT,
    CONSTRAINT fk_events FOREIGN KEY(event_id_fk) REFERENCES events(event_id)
);
