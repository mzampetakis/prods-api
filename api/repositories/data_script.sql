TRUNCATE `products`;
TRUNCATE `categories`;

INSERT INTO `categories` (`id`, `title`, `sort`, `image_url`)
VALUES
	(1,'Laptops',1,'https://category1.image'),
	(2,'Monitors',2,'https://category2.image'),
	(3,'Keyboards',3,'https://category3.image'),
	(4,'Mice',4,'https://category4.image'),
	(5,'USB Sticks',5,'https://category5.image');

INSERT INTO `products` (`id`, `category_id`, `title`, `image_url`, `price`, `description`)
VALUES
	(1,1,'Laptop 15','https://product1.image',150000,'Some Laptop1 Description'),
	(2,1,'Laptop 16','https://product2.image',160000,'Some other Description fro product #2'),
	(3,1,'Laptop17','https://product3.image',170000,'Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. Some realy big Description. '),
	(4,1,'Laptop 13','https://product4.image',130000,NULL),
	(5,1,'Laptop 20','https://product5.image',200000,NULL),
	(6,2,'Ultrasharp 21','https://product6.image',21000,NULL),
	(7,2,'Ultrasharp 24','https://product7.image',24000,'Description of a really good monitor'),
	(8,2,'Ultrasharp 27','https://product8.image',27000,'VFM monitor\n'),
	(9,2,'Ultrasharp 30','https://product9.image',30000,NULL),
	(10,2,'OLED 21','https://product10.image',31000,NULL),
	(11,2,'OLED 24','https://product11.image',34000,'Description of a really good monitor'),
	(12,2,'OLED 27','https://product12.image',37000,'VFM monitor\n'),
	(13,2,'OLED 30','https://product13.image',40000,NULL),
	(14,5,'1GB','https://product10.image',1050,'The biggest flash drive ever!');