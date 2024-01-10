CREATE OR REPLACE FUNCTION get_future_tours_by_country(
    country_name TEXT
) RETURNS TABLE (
    tour_id INTEGER,
    tour_name TEXT,
    start_date DATE,
    end_date DATE,
    price MONEY
) AS $$
BEGIN
    RETURN QUERY
    SELECT
        t.id AS tour_id,
        t.name AS tour_name,
        t.start_date,
        t.end_date,
        t.price
    FROM
        Tour t
    INNER JOIN
        Trip tr ON t.id = tr.tour_id
    INNER JOIN
        Country c ON tr.dst_country_id = c.id
    WHERE
        c.name = country_name
        AND t.start_date > CURRENT_DATE; -- Assuming you want future tours

END;
$$ LANGUAGE plpgsql;

create or replace procedure book_hotel_room(
    tourist_id integer,
    hotel_id integer,
    room_id integer,
    start_date date,
    end_date date
) as $$
begin
    insert into Tourist_hotel(tourist_id, hotel_id, room_id, start_date, end_date)
    values (tourist_id, hotel_id, room_id, start_date, end_date);
end;
$$ language plpgsql;

create or replace procedure book_tour(
    tourist_id integer,
    tour_id integer
) as $$
begin
    insert into Tourist_tour(tourist_id, tour_id)
    values (tourist_id, tour_id);
end;
$$ language plpgsql;

create or replace procedure create_trip(
    tour_id integer,
    src_country text,
    dst_country text,
    transport_type text,
    datetime_start timestamp,
    datetime_end timestamp,
    price money
) as $$
begin 
    insert into Trip(tour_id, src_country_id, dst_country_id, transport_type, datetime_start, datetime_end, price)
    values (tour_id, (select id from Country where name = src_country), (select id from Country where name = dst_country), transport_type, datetime_start, datetime_end, price);
end;
$$ language plpgsql;

create or replace function get_free_rooms_by_hotel(
    arg_hotel_id integer,
    arg_start_date date,
    arg_end_date date
) returns table (
    room_id integer,
    places smallint,
    price money
) as $$
begin
    return query
    select
        r.room_id,
        r.places_number,
        r.price_per_night
    from
        hotel_room r
    where
        r.room_id not in (
            select
                th.room_id
            from
                tourist_hotel th
            where
                th.hotel_id = arg_hotel_id
                and (
                    (th.start_date <= arg_start_date and th.end_date >= arg_start_date)
                    or (th.start_date <= arg_end_date and th.end_date >= arg_end_date)
                    or (th.start_date >= arg_start_date and th.end_date <= arg_end_date)
                )
        ) and r.hotel_id = arg_hotel_id;
end;
$$ language plpgsql;

create or replace function get_info_for_visa(
    arg_tourist_id integer,
    arg_tour_id integer
) returns table (
    country_id integer
) as $$
begin
    return query SELECT 
        Trip.dst_country_id 
    FROM Trip
    WHERE Trip.tour_id = arg_tour_id
    AND (Trip.dst_country_id NOT IN (
        SELECT Visa.destination_country_id FROM Visa 
        WHERE Visa.tourist_id = arg_tourist_id and Visa.expiry_date >= (select end_date from tour where tour.id = arg_tour_id) and Visa.issue_date <= (select start_date from tour where tour.id = arg_tour_id)
        )
    and Trip.dst_country_id != (select citizenship_id from tourist where tourist.id = arg_tourist_id)
        );
end;
$$ language plpgsql;



-- create or replace procedure issue_visa(
--     arg_tourist_id integer,
--     arg_tour_id integer
-- ) as $$
-- DECLARE
--     cntry_id integer;
--     country_name text;
-- begin
--     select country_id into cntry_id from get_info_for_visa(arg_tourist_id, arg_tour_id);
--     if cntry_id is null then
--         raise notice 'null';
--         return;
--     end if;
--     select "name" into country_name from country where id = cntry_id;
--     if country_name is null then
--         raise notice 'null';
--         return;
--     end if;


--     insert into Visa
--     ("name", destination_country_id, tourist_id, issue_date, expiry_date)
--     values 
--     (country_name, 
--     cntry_id, 
--     arg_tourist_id, (select start_date from tour where tour.id = arg_tour_id), 
--     (select end_date from tour where tour.id = arg_tour_id));
-- end;
-- $$ language plpgsql;

CREATE OR REPLACE PROCEDURE issue_visa(
    arg_tourist_id INTEGER,
    arg_tour_id INTEGER
) AS $$
DECLARE
    cntry_id INTEGER;
    country_name TEXT;
    country_cursor CURSOR FOR
        SELECT country_id
        FROM get_info_for_visa(arg_tourist_id, arg_tour_id);
BEGIN
    OPEN country_cursor;
    LOOP
        FETCH country_cursor INTO cntry_id;
        EXIT WHEN NOT FOUND;

        IF cntry_id IS NULL THEN
            RAISE NOTICE 'null';
            CONTINUE; -- Skip to the next iteration if country_id is null
        END IF;

        SELECT "name" INTO country_name FROM country WHERE id = cntry_id;
        IF country_name IS NULL THEN
            RAISE NOTICE 'null';
            CONTINUE; -- Skip to the next iteration if country_name is null
        END IF;

        INSERT INTO Visa
        ("name", destination_country_id, tourist_id, issue_date, expiry_date)
        VALUES 
        (country_name, 
        cntry_id, 
        arg_tourist_id, (SELECT start_date FROM tour WHERE id = arg_tour_id), 
        (SELECT end_date FROM tour WHERE id = arg_tour_id));
    END LOOP;

    CLOSE country_cursor;
END;
$$ LANGUAGE plpgsql;
