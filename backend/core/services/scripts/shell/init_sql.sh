# hard symbolic link
# user sql
ln ../../user-service/sql/init.sql ../sql/010-init.sql
ln ../../user-service/sql/data.sql ../sql/011-data.sql

# follow sql
ln ../../follow-service/sql/init.sql ../sql/020-init.sql
ln ../../follow-service/sql/data.sql ../sql/021-data.sql

# post sql
ln ../../post-service/sql/init.sql ../sql/030-init.sql
ln ../../post-service/sql/data.sql ../sql/031-data.sql

# comment sql
ln ../../comment-service/sql/init.sql ../sql/040-init.sql
ln ../../comment-service/sql/data.sql ../sql/041-data.sql