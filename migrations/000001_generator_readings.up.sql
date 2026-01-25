CREATE TABLE generator_readings (
    id                SERIAL    PRIMARY KEY,
    gasoline          INTEGER   NOT NULL,
    generator_voltage INTEGER   NOT NULL,
    network_voltage   INTEGER   NOT NULL,
    time              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

