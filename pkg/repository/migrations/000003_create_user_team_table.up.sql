CREATE TABLE IF NOT EXISTS user_team(
    user_id integer not null,
    team_id integer not null,
    PRIMARY KEY (user_id, team_id),
    FOREIGN KEY (user_id) REFERENCES "user" (id) ON DELETE CASCADE,
    FOREIGN KEY (team_id) REFERENCES "team" (id) ON DELETE CASCADE
);