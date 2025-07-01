-- Active: 1750082338800@@127.0.0.1@5432@postgres
INSERT INTO
    users (
        id,
        name,
        phone_number,
        profile_image,
        email,
        password,
        role
    )
VALUES (
        'u-001',
        'Hosea Zen',
        '081234567890',
        'profile1.jpg',
        'hosea@gmail.com',
        '12345',
        'admin'
    ),
    (
        'u-002',
        'Bayu Adi',
        '081298765432',
        'profile2.jpg',
        'bayu@gmail.com',
        '22345',
        'users'
    ),
    (
        'u-003',
        'Mas Yasirstyle',
        '081111111111',
        'profile3.jpg',
        'yasir@gmail.com',
        '1224345',
        'users'
    ),
    (
        'u-004',
        'Mas Nanda',
        '08111111232',
        'profile4.jpg',
        'nanda@gmail.com',
        '1224345',
        'users'
    ),
    (
        'u-005',
        'Mas asa',
        '0811112',
        'profile5.jpg',
        'asa@gmail.com',
        '12243456',
        'users'
    );

INSERT INTO
    genres (id, name)
VALUES ('g-001', 'Action'),
    ('g-002', 'Comedy'),
    ('g-003', 'Drama'),
    ('g-004', 'Horor'),
    ('g-005', 'fantasy'),
    ('g-006', 'Comedy');

INSERT INTO
    directors (id, name)
VALUES ('d-001', 'Christopher Nolan'),
    ('d-002', 'Quentin Tarantino');

INSERT INTO
    actors (id, name)
VALUES ('a-001', 'Leonardo DiCaprio'),
    ('a-002', 'Brad Pitt'),
    ('a-003', 'Cillian Murphy');

INSERT INTO
    movies (
        id,
        title,
        synopsis,
        background,
        poster,
        release_date,
        duration,
        price
    )
VALUES (
        'm-001',
        'Inception',
        'A mind-bending thriller.',
        'bg1.jpg',
        'poster1.jpg',
        '2010-07-16 00:00:00',
        148,
        50000
    ),
    (
        'm-002',
        'Once Upon a Time in Hollywood',
        'An actor and stunt double...',
        'bg2.jpg',
        'poster2.jpg',
        '2019-07-26 00:00:00',
        161,
        55000
    ),
    (
        'm-003',
        'Once  in Hollywood',
        'An actor and stunt double...',
        'bg2.jpg',
        'poster2.jpg',
        '2019-07-26 00:00:00',
        161,
        55000
    ),
    (
        'm-004',
        'Osean',
        'An actor and stunt double...',
        'bg2.jpg',
        'poster2.jpg',
        '2019-07-26 00:00:00',
        161,
        55000
    ),
    (
        'm-005',
        'O Hollywood',
        'An actor and stunt double...',
        'bg2.jpg',
        'poster2.jpg',
        '2019-07-26 00:00:00',
        161,
        55000
    );

INSERT INTO
    movie_director (id, movie_id, director_id)
VALUES ('md-001', 'm-001', 'd-001'),
    ('md-002', 'm-002', 'd-002');

INSERT INTO
    movie_actors (id, movie_id, actor_id)
VALUES ('ma-001', 'm-001', 'a-003'),
    ('ma-002', 'm-002', 'a-001'),
    ('ma-003', 'm-002', 'a-002');

INSERT INTO
    cinema (id, name)
VALUES ('c-001', 'Cinepolis'),
    ('c-002', 'CGV'),
    ('c-003', 'XX1'),
    ('c-004', 'Icon Walk XXI');

INSERT INTO
    payment_method (id, name)
VALUES ('p-001', 'Go-Pay'),
    ('p-002', 'Dana'),
    ('p-003', 'BRI'),
    ('p-004', 'PayPal');

INSERT INTO
    transactions (
        id,
        time,
        date,
        id_cinema,
        id_payment_method,
        price_total,
        user_id,
        movie_id
    )
VALUES (
        't-006',
        '2025-06-19 15:30:00',
        '2025-06-19 00:00:00',
        'c-001',
        'p-001',
        100000,
        'u-002',
        'm-001'
    ),
    (
        't-002',
        '2025-06-19 15:30:00',
        '2025-06-19 00:00:00',
        'c-002',
        'p-004',
        100000,
        'u-002',
        'm-001'
    ),
    (
        't-004',
        '2025-06-19 15:30:00',
        '2025-06-19 00:00:00',
        'c-003',
        'p-001',
        100000,
        'u-001',
        'm-003'
    );

DELETE FROM transaction_detail WHERE costumer_name = 'hosea';

INSERT INTO
    transaction_detail (
        id,
        id_transaction,
        costumer_name,
        costumer_phone,
        seat
    )
VALUES (
        'td-001',
        't-002',
        'hosea',
        '081411110',
        'A1'
    ),
    (
        'td-002',
        't-006',
        'hosea',
        '081411110',
        'A2'
    ),
    (
        'td-003',
        't-004',
        'hosea',
        '081411110',
        'A2'
    );

INSERT INTO
    history_transaction (
        id,
        transaction_id,
        status,
        updated_by,
        note
    )
VALUES (
        'ht-001',
        't-006',
        'paid',
        'admin',
        'Payment confirmed by admin'
    ),
    (
        'ht-002',
        't-002',
        'paid',
        'admin',
        'Payment confirmed by admin'
    ),
    (
        'ht-003',
        't-004',
        'paid',
        'admin',
        'Payment confirmed by admin'
    );

SELECT * FROM transactions;

SELECT * FROM transaction_detail;

SELECT * FROM history_transaction;

-- Update data
INSERT INTO
    users (
        id,
        name,
        phone_number,
        profile_image,
        email,
        password,
        role
    )
VALUES (
        'u-006',
        'Dina Ayu',
        '081234511111',
        'profile6.jpg',
        'dina@gmail.com',
        'pwd6',
        'users'
    ),
    (
        'u-007',
        'Riko Maulana',
        '082234567891',
        'profile7.jpg',
        'riko@gmail.com',
        'pwd7',
        'users'
    ),
    (
        'u-008',
        'Sinta Dewi',
        '083334567892',
        'profile8.jpg',
        'sinta@gmail.com',
        'pwd8',
        'users'
    ),
    (
        'u-009',
        'Taufik Hidayat',
        '084434567893',
        'profile9.jpg',
        'taufik@gmail.com',
        'pwd9',
        'users'
    ),
    (
        'u-010',
        'Yuni Permata',
        '085534567894',
        'profile10.jpg',
        'yuni@gmail.com',
        'pwd10',
        'users'
    );

INSERT INTO
    genres (id, name)
VALUES ('g-007', 'Sci-Fi'),
    ('g-008', 'Romance'),
    ('g-009', 'Adventure'),
    ('g-010', 'Thriller');

INSERT INTO
    directors (id, name)
VALUES ('d-003', 'James Cameron'),
    ('d-004', 'Steven Spielberg'),
    ('d-005', 'Patty Jenkins'),
    ('d-006', 'Zack Snyder'),
    ('d-007', 'Greta Gerwig'),
    ('d-008', 'Denis Villeneuve');

INSERT INTO
    actors (id, name)
VALUES ('a-004', 'Tom Holland'),
    ('a-005', 'Emma Watson'),
    ('a-006', 'Robert Downey Jr.'),
    ('a-007', 'Gal Gadot'),
    ('a-008', 'Tom Hanks'),
    ('a-009', 'Scarlett Johansson'),
    ('a-010', 'Chris Hemsworth');

INSERT INTO
    movies (
        id,
        title,
        synopsis,
        background,
        poster,
        release_date,
        duration,
        price
    )
VALUES (
        'm-006',
        'Interstellar',
        'A space exploration movie.',
        'bg3.jpg',
        'poster3.jpg',
        '2014-11-07',
        169,
        60000
    ),
    (
        'm-007',
        'Titanic',
        'A romantic disaster film.',
        'bg4.jpg',
        'poster4.jpg',
        '1997-12-19',
        195,
        45000
    ),
    (
        'm-008',
        'Avengers: Endgame',
        'Superheroes fight to save the universe.',
        'bg5.jpg',
        'poster5.jpg',
        '2019-04-26',
        181,
        70000
    ),
    (
        'm-009',
        'Wonder Woman',
        'A warrior princess fights for peace.',
        'bg6.jpg',
        'poster6.jpg',
        '2017-06-02',
        141,
        50000
    ),
    (
        'm-010',
        'Barbie',
        'A fantasy-comedy adventure.',
        'bg7.jpg',
        'poster7.jpg',
        '2023-07-21',
        114,
        55000
    );

INSERT INTO
    movie_director (id, movie_id, director_id)
VALUES ('md-003', 'm-003', 'd-003'),
    ('md-004', 'm-004', 'd-004'),
    ('md-005', 'm-005', 'd-005'),
    ('md-006', 'm-006', 'd-006'),
    ('md-007', 'm-007', 'd-007'),
    ('md-008', 'm-008', 'd-001'),
    ('md-009', 'm-009', 'd-005'),
    ('md-010', 'm-010', 'd-007');

INSERT INTO
    movie_actors (id, movie_id, actor_id)
VALUES ('ma-004', 'm-003', 'a-004'),
    ('ma-005', 'm-004', 'a-005'),
    ('ma-006', 'm-005', 'a-006'),
    ('ma-007', 'm-006', 'a-007'),
    ('ma-008', 'm-007', 'a-008'),
    ('ma-009', 'm-008', 'a-009'),
    ('ma-010', 'm-009', 'a-010');

INSERT INTO
    cinema (id, name)
VALUES ('c-005', 'Grand Cinema'),
    ('c-006', 'Blitz Megaplex');

INSERT INTO
    payment_method (id, name)
VALUES ('p-005', 'OVO'),
    ('p-006', 'ShopeePay');

INSERT INTO
    transactions (
        id,
        time,
        date,
        id_cinema,
        id_payment_method,
        price_total,
        user_id,
        movie_id
    )
VALUES (
        't-005',
        '2025-06-20 17:00:00',
        '2025-06-20',
        'c-001',
        'p-002',
        55000,
        'u-006',
        'm-004'
    ),
    (
        't-007',
        '2025-06-21 19:00:00',
        '2025-06-21',
        'c-002',
        'p-003',
        45000,
        'u-007',
        'm-005'
    ),
    (
        't-008',
        '2025-06-22 20:00:00',
        '2025-06-22',
        'c-003',
        'p-001',
        60000,
        'u-008',
        'm-006'
    ),
    (
        't-009',
        '2025-06-22 21:00:00',
        '2025-06-22',
        'c-004',
        'p-004',
        70000,
        'u-009',
        'm-007'
    ),
    (
        't-010',
        '2025-06-23 22:00:00',
        '2025-06-23',
        'c-001',
        'p-005',
        55000,
        'u-010',
        'm-008'
    );

INSERT INTO
    transaction_detail (
        id,
        id_transaction,
        costumer_name,
        costumer_phone,
        seat
    )
VALUES (
        'td-004',
        't-005',
        'dina',
        '081234511111',
        'A3'
    ),
    (
        'td-005',
        't-007',
        'riko',
        '082234567891',
        'B1'
    ),
    (
        'td-006',
        't-008',
        'sinta',
        '083334567892',
        'B2'
    ),
    (
        'td-007',
        't-009',
        'taufik',
        '084434567893',
        'C1'
    ),
    (
        'td-008',
        't-010',
        'yuni',
        '085534567894',
        'C2'
    );

INSERT INTO
    history_transaction (
        id,
        transaction_id,
        status,
        updated_by,
        note
    )
VALUES (
        'ht-004',
        't-005',
        'paid',
        'admin',
        'Confirmed'
    ),
    (
        'ht-005',
        't-007',
        'paid',
        'admin',
        'Confirmed'
    ),
    (
        'ht-006',
        't-008',
        'paid',
        'admin',
        'Confirmed'
    ),
    (
        'ht-007',
        't-009',
        'paid',
        'admin',
        'Confirmed'
    ),
    (
        'ht-008',
        't-010',
        'paid',
        'admin',
        'Confirmed'
    );