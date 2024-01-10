create index tourist_hotel_date_idx on tourist_hotel(start_date, end_date);

create index trip_date_idx on trip(datetime_start, datetime_end);

create index trip_tourid_idx on trip(tour_id);

create index trip_dst_cntr_idx on trip(dst_country_id);

create index visa_touristid_idx on visa(tourist_id);

create index tourist_hotel_hotelid_idx on tourist_hotel(hotel_id);