LOAD DATA INFILE '/Users/zhuhaowei/Desktop/Cabin/yagoo/resources/wukong_test.csv' 
INTO TABLE wukong_data
FIELDS TERMINATED BY ',' 
ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 1 ROWS
(dir, text, token, url, alt_id);