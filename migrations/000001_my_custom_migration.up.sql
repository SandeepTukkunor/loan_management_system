CREATE TABLE Users (
    UserID UUID PRIMARY KEY,
    Email VARCHAR(255) NOT NULL,
    Password VARCHAR(255) NOT NULL,
    Mobile BIGINT,
    IsActive BOOLEAN NOT NULL DEFAULT false,
    IsStaff BOOLEAN NOT NULL DEFAULT false,
    IsEmailVerified BOOLEAN NOT NULL DEFAULT false,
    IsMobileVerified BOOLEAN NOT NULL DEFAULT false,
    InfoID UUID
);