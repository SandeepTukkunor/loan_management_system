
CREATE TABLE Users (
    userId uuid PRIMARY KEY,
    Email varchar(255) NOT NULL,
    Password varchar(255) NOT NULL,
    mobile bigint,
    isActive boolean NOT NULL DEFAULT false,
    isStaff boolean NOT NULL DEFAULT false,
    isEmailEverified boolean NOT NULL DEFAULT false,
    infoID uuid
);
