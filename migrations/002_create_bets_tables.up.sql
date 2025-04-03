CREATE TABLE IF NOT EXISTS bets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    event_id UUID NOT NULL,  -- Links to a sports event - events table TODO
    amount NUMERIC(10,2) NOT NULL CHECK (amount > 0),
    odds NUMERIC(5,2) NOT NULL CHECK (odds > 0), -- Betting odds (1.50, 2.75) - these are to be mapped from an external api
    bet_type TEXT NOT NULL CHECK (bet_type IN ('SINGLE', 'MULTI', 'SYSTEM')),
    status TEXT NOT NULL DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'WON', 'LOST', 'CANCELLED')),
    outcome TEXT CHECK (outcome IN ('WIN', 'LOSE', 'VOID')),
    placed_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),

    -- Foreign Key Constraints - these were suppose to be the same as .proto descriptions
    CONSTRAINT fk_bet_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_bet_event FOREIGN KEY (event_id) REFERENCES events (id) ON DELETE CASCADE
);

-- Indexes for performance - still being tested
CREATE INDEX idx_bets_user_id ON bets (user_id);
CREATE INDEX idx_bets_event_id ON bets (event_id);
CREATE INDEX idx_bets_status ON bets (status);