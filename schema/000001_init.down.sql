DROP TABLE IF EXISTS companies;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS suppliers;
DROP TABLE IF EXISTS purchased_materials_archive;
DROP TABLE IF EXISTS planning_materials_archive;
DROP TABLE IF EXISTS purchased_materials;
DROP TABLE IF EXISTS planning_materials;
DROP TABLE IF EXISTS warehouses;
DROP TABLE IF EXISTS refresh_tokens;
DROP TABLE IF EXISTS sections;

DELETE
FROM "users"
WHERE "username" = 'admin';