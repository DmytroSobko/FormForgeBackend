CREATE TABLE athletes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    athlete_type TEXT NOT NULL,
    name TEXT NOT NULL,

    strength DOUBLE PRECISION NOT NULL CHECK (strength BETWEEN 0 AND 100),
    endurance DOUBLE PRECISION NOT NULL CHECK (endurance BETWEEN 0 AND 100),
    mobility DOUBLE PRECISION NOT NULL CHECK (mobility BETWEEN 0 AND 100),

    fatigue DOUBLE PRECISION NOT NULL CHECK (fatigue >= 0),
    max_fatigue DOUBLE PRECISION NOT NULL CHECK (max_fatigue > 0),

    week INT NOT NULL DEFAULT 1,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_athletes_type ON athletes(athlete_type);