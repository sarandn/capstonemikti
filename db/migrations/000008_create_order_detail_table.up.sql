CREATE TABLE OrderDetail (
    order_detail_id SERIAL PRIMARY KEY,
    order_id_fk INT,
    ticket_id_fk INT,
    quantity INT,
    subtotal INT,
    CONSTRAINT fk_orders FOREIGN KEY (order_id_fk) REFERENCES orders(order_id),
    CONSTRAINT fk_tickets FOREIGN KEY (ticket_id_fk) REFERENCES Ticket(ticket_id)
);
