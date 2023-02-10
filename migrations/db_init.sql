-- create new database
CREATE DATABASE `freefolks-fc`;

-- Enum Type ui_body_size
CREATE TYPE body_size AS ENUM('XXS', 'XS', 'S', 'M', 'L', 'XL', 'XXL', 'XXXL');

-- user_information
CREATE TABLE user_information(
  ui_id SERIAL NOT NULL,
  ui_user_type SMALLINT,
  ui_customer_number VARCHAR(255),
  ui_first_name VARCHAR(255),
  ui_last_name VARCHAR(255),
  ui_email VARCHAR(255),
  ui_password VARCHAR(255),
  ui_mobile_number VARCHAR(255),
  ui_occupation INT,
  ui_date_of_birth DATE,
  ui_gender VARCHAR(255),
  ui_photo_profile VARCHAR(255),
  ui_address TEXT,
  ui_city VARCHAR(255),
  ui_postal_code VARCHAR(255),
  ui_body_size body_size,
  ui_activation_code VARCHAR(255),
  ui_email_status BOOLEAN,
  ui_verified_at TIMESTAMP,
  ui_created_at TIMESTAMP DEFAULT NOW(),
  ui_created_by INT,
  ui_updated_at TIMESTAMP,
  ui_updated_by INT,
  PRIMARY KEY (ui_id)
);

-- Comments
COMMENT ON COLUMN user_information.ui_occupation IS 'Reference to mr_occupation';

-- Function for ui_updated_at 
CREATE FUNCTION trigger_ui_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.ui_updated_at = NOW();
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_ui_timestamp
BEFORE
UPDATE ON user_information
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

-- =============================================================================================

-- preferred_position
CREATE TABLE preferred_position(
  pp_id SERIAL NOT NULL,
  pp_ui_id INT,
  pp_position INT,
  pp_created_at TIMESTAMP DEFAULT NOW(),
  pp_created_by INT,
  pp_updated_at TIMESTAMP DEFAULT NOW(),
  pp_updated_by INT,
  PRIMARY KEY (pp_id),
  FOREIGN KEY (pp_ui_id) REFERENCES user_information(ui_id)
);

-- Comments
COMMENT ON COLUMN preferred_position.pp_position IS 'Reference to mr_position';

-- =============================================================================================

-- game_data
CREATE TABLE game_data(
  gd_id SERIAL NOT NULL,
  gd_game_number VARCHAR(255),
  gd_venue_name VARCHAR(255),
  gd_venue_address VARCHAR(255),
  gd_map_url VARCHAR(255),
  gd_time TIMESTAMP,
  gd_goalkeeper_quota INT,
  gd_outfield_quota INT,
  gd_goalkeeper_price DOUBLE PRECISION,
  gd_outfield_price DOUBLE PRECISION,
  gd_total_cost DOUBLE PRECISION,
  gd_notes TEXT,
  gd_status BOOLEAN,
  gd_created_at TIMESTAMP DEFAULT NOW(),
  gd_created_by INT,
  gd_updated_at TIMESTAMP,
  gd_updated_by INT,
  PRIMARY KEY (gd_id)
);

-- =============================================================================================

-- game_information
CREATE TABLE game_information(
  gi_id SERIAL NOT NULL,
  gi_gd_id INT,
  gi_type INT,
  gi_description TEXT,
  gi_created_at TIMESTAMP DEFAULT NOW(),
  gi_created_by INT,
  gi_updated_at TIMESTAMP,
  gi_updated_by INT,
  PRIMARY KEY (gi_id),
  FOREIGN KEY (gi_gd_id) REFERENCES game_data(gd_id)
);

-- =============================================================================================

-- game_galleries
CREATE TABLE game_galleries(
  ggs_id SERIAL NOT NULL,
  ggs_gd_id INT,
  ggs_image_url TEXT,
  ggs_alt_image VARCHAR(255),
  ggs_created_at TIMESTAMP DEFAULT NOW(),
  ggs_created_by INT,
  ggs_updated_at TIMESTAMP,
  ggs_updated_by INT,
  PRIMARY KEY (ggs_id),
  FOREIGN KEY (ggs_gd_id) REFERENCES game_data(gd_id)
);

-- =============================================================================================

-- game_costs
CREATE TABLE game_costs(
  gcs_id SERIAL NOT NULL,
  gcs_gd_id INT,
  gcs_description TEXT,
  gcs_cost DOUBLE PRECISION,
  gcs_created_at TIMESTAMP DEFAULT NOW(),
  gcs_created_by INT,
  gcs_updated_at TIMESTAMP,
  gcs_updated_by INT,
  PRIMARY KEY (gcs_id),
  FOREIGN KEY (gcs_gd_id) REFERENCES game_data(gd_id)
);

-- =============================================================================================

-- game_registration
CREATE TABLE game_registration(
  gr_id SERIAL NOT NULL,
  gr_gd_id INT,
  gr_ui_id INT,
  gr_is_outfield BOOLEAN,
  gr_amount DOUBLE PRECISION,
  gr_transaction_number VARCHAR(255),
  gr_created_at TIMESTAMP DEFAULT NOW(),
  gr_created_by INT,
  gr_updated_at TIMESTAMP,
  gr_updated_by INT,
  PRIMARY KEY (gr_id),
  FOREIGN KEY (gr_gd_id) REFERENCES game_data(gd_id),
  FOREIGN KEY (gr_ui_id) REFERENCES user_information(ui_id)
);

-- =============================================================================================

-- game_registered_player
CREATE TABLE game_registered_player(
  grp_id SERIAL NOT NULL,
  grp_gd_id INT,
  grp_ui_id INT,
  grp_is_outfield BOOLEAN,
  grp_amount_paid DOUBLE PRECISION,
  grp_paid_at TIMESTAMP,
  grp_transaction_number VARCHAR(255),
  grp_created_at TIMESTAMP DEFAULT NOW(),
  grp_created_by INT,
  grp_updated_at TIMESTAMP,
  grp_updated_by INT,
  PRIMARY KEY (grp_id),
  FOREIGN KEY (grp_gd_id) REFERENCES game_data(gd_id),
  FOREIGN KEY (grp_ui_id) REFERENCES user_information(ui_id)
);
