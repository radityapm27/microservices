
INSERT INTO "public"."cart" ("cart_id", "user_id", "product_id", "quantity") VALUES
('13bc7f0f-04d6-485d-a3c0-c305fbbf54e1', 'test2@gmail.com', '11a3e38e-8a96-4eb7-901a-d0925289f10a', 30),
('5618d7c8-d4d3-4ac9-a7ff-f09241c9e637', 'test1@gmail.com', 'db55f33f-e6f0-4715-b6f5-715556ca5336', 20),
('5bc8cb2f-fc27-423f-96fe-9237f0b3c838', 'test1@gmail.com', '3415c41a-2150-49c7-bc53-f1aed6d6a69d', 10),
('d1ea3e05-52e7-4494-8ee5-5ab7032de7a3', 'test2@gmail.com', '89ea8626-a747-4279-bb99-5130b1eb4f89', 40);

INSERT INTO "public"."catalog" ("product_id", "product_name", "product_desc", "price", "stock", "created_time", "created_by", "updated_time", "updated_by") VALUES
('11a3e38e-8a96-4eb7-901a-d0925289f10a', 'Product Name 3', 'Product Desc 3', 3000, 300, '2022-04-03 20:05:35.804375+07', 'system', '2022-04-03 20:05:35.804375+07', 'system'),
('3415c41a-2150-49c7-bc53-f1aed6d6a69d', 'Product Name 1', 'Product Desc 1', 1000, 100, '2022-04-03 20:05:35.804375+07', 'system', '2022-04-03 20:05:35.804375+07', 'system'),
('89ea8626-a747-4279-bb99-5130b1eb4f89', 'Product Name 4', 'Product Desc 4', 4000, 400, '2022-04-03 20:05:35.804375+07', 'system', '2022-04-03 20:05:35.804375+07', 'system'),
('db55f33f-e6f0-4715-b6f5-715556ca5336', 'Product Name 2', 'Product Desc 2', 2000, 200, '2022-04-03 20:05:35.804375+07', 'system', '2022-04-03 20:05:35.804375+07', 'system');

INSERT INTO "public"."users" ("user_id", "user_name", "user_location") VALUES
('test1@gmail.com', 'User Test 1', 'Jakarta'),
('test2@gmail.com', 'User Test 2', 'Jakarta');

ALTER TABLE "public"."cart" ADD FOREIGN KEY ("product_id") REFERENCES "public"."catalog"("product_id");
ALTER TABLE "public"."cart" ADD FOREIGN KEY ("user_id") REFERENCES "public"."users"("user_id");
