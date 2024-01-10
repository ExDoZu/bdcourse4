
CREATE OR REPLACE FUNCTION check_room_availability()
RETURNS TRIGGER as $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM Hotel_room
        WHERE hotel_id = NEW.hotel_id
        AND room_id = NEW.room_id
    ) THEN
        RAISE EXCEPTION 'Selected room is not exists in the specified hotel.';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM Tourist_hotel
        WHERE hotel_id = NEW.hotel_id
        AND room_id = NEW.room_id
        AND (
            (NEW.start_date between start_date AND end_date)
            OR (NEW.end_date between start_date AND end_date)
            or (NEW.start_date <= start_date AND NEW.end_date >= end_date)
        )
    ) THEN
        RAISE EXCEPTION 'Selected room is already booked for the specified dates.';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_check_room_availability
BEFORE INSERT OR UPDATE ON Tourist_hotel
FOR EACH ROW
EXECUTE FUNCTION check_room_availability();


CREATE OR REPLACE FUNCTION check_visa_validity()
RETURNS TRIGGER AS $$
DECLARE
    visa_not_valid BOOLEAN;
BEGIN
    SELECT COUNT(*) > 0 INTO visa_not_valid
    FROM get_info_for_visa(NEW.tourist_id, NEW.tour_id);

    IF visa_not_valid THEN
        RAISE EXCEPTION 'The tourist does not have a valid visa for the tour';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER check_visa_validity_trigger
BEFORE INSERT OR UPDATE ON Tourist_tour
FOR EACH ROW EXECUTE FUNCTION check_visa_validity();

CREATE OR REPLACE FUNCTION update_tour_satisfaction()
RETURNS TRIGGER AS $$
BEGIN
    UPDATE Tour
    SET satisfaction_level = (
        SELECT AVG(rating)
        FROM Tour_review tr
        WHERE tr.tour_id = NEW.tour_id
    )
    WHERE id = NEW.tour_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_tour_satisfaction_trigger
AFTER INSERT ON Tour_review
FOR EACH ROW EXECUTE FUNCTION update_tour_satisfaction();





