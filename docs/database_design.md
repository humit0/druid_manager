# Database Design

## datasource
|column_name|column_type|extra|
|---|---|---|
|id|bigint|auto_increment|
|name|varchar(255)|NOT NULL|
|owner|varchar(80)|DEFAULT NULL|
|interval|bigint||
|monitor|tinyint(1)|DEFAULT 0|

## datasource_column
|column_name|column_type|extra|
|---|---|---|
|id|bigint|auto_increment|
|datasource_id|bigint|KEY|
|column_name|varchar(255)|NOT NULL|
|column_type|enum(dimension, metric)||
|is_count|bool|DEFAULT FALSE|

## datasource_meta
|column_name|column_type|extra|
|---|---|---|
|id|bigint|auto_increment|
|datasource_id|bigint|KEY|
|date_id|timestamp|KEY|
|row_cnt|bigint||
|segment_cnt|bigint||
|byte_size|bigint||
|rollup_ratio|decimal||

## datasource_column_meta
|column_name|column_type|extra|
|---|---|---|
|id|bigint|auto_increment|
|column_id|bigint|KEY|
|date_id|timestamp|KEY|
|cardinarity|bigint|DEFAULT 0|
