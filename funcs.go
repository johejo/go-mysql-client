package gomysqlclient

var mysqlFuncs = []string{
	"ABS()",
	"ACOS()",
	"ADDDATE()",
	"ADDTIME()",
	"AES_DECRYPT()",
	"AES_ENCRYPT()",
	"AND",
	"&&",
	"ANY_VALUE()",
	"Area()",
	"AsBinary()",
	"AsWKB()",
	"ASCII()",
	"ASIN()",
	"=",
	":=",
	"AsText()",
	"AsWKT()",
	"ASYMMETRIC_DECRYPT()",
	"ASYMMETRIC_DERIVE()",
	"ASYMMETRIC_ENCRYPT()",
	"ASYMMETRIC_SIGN()",
	"ASYMMETRIC_VERIFY()",
	"ATAN()",
	"ATAN2()",
	"ATAN()",
	"AVG()",
	"BENCHMARK()",
	"BETWEEN ... AND ...",
	"BIN()",
	"BIN_TO_UUID()",
	"BINARY",
	"BIT_AND()",
	"BIT_COUNT()",
	"BIT_LENGTH()",
	"BIT_OR()",
	"BIT_XOR()",
	"&",
	"~",
	"|",
	"^",
	"Buffer()",
	"CAN_ACCESS_COLUMN()",
	"CAN_ACCESS_DATABASE()",
	"CAN_ACCESS_TABLE()",
	"CAN_ACCESS_VIEW()",
	"CASE",
	"CAST()",
	"CEIL()",
	"CEILING()",
	"Centroid()",
	"CHAR()",
	"CHAR_LENGTH()",
	"CHARACTER_LENGTH()",
	"CHARSET()",
	"COALESCE()",
	"COERCIBILITY()",
	"COLLATION()",
	"COMPRESS()",
	"CONCAT()",
	"CONCAT_WS()",
	"CONNECTION_ID()",
	"Contains()",
	"CONV()",
	"CONVERT()",
	"CONVERT_TZ()",
	"ConvexHull()",
	"COS()",
	"COT()",
	"COUNT()",
	"COUNT(DISTINCT)",
	"CRC32()",
	"CREATE_ASYMMETRIC_PRIV_KEY()",
	"CREATE_ASYMMETRIC_PUB_KEY()",
	"CREATE_DH_PARAMETERS()",
	"CREATE_DIGEST()",
	"Crosses()",
	"CUME_DIST()",
	"CURDATE()",
	"CURRENT_DATE()",
	"CURRENT_DATE",
	"CURRENT_ROLE()",
	"CURRENT_TIME()",
	"CURRENT_TIME",
	"CURRENT_TIMESTAMP()",
	"CURRENT_TIMESTAMP",
	"CURRENT_USER()",
	"CURRENT_USER",
	"CURTIME()",
	"DATABASE()",
	"DATE()",
	"DATE_ADD()",
	"DATE_FORMAT()",
	"DATE_SUB()",
	"DATEDIFF()",
	"DAY()",
	"DAYNAME()",
	"DAYOFMONTH()",
	"DAYOFWEEK()",
	"DAYOFYEAR()",
	"DECODE()",
	"DEFAULT()",
	"DEGREES()",
	"DENSE_RANK()",
	"DES_DECRYPT()",
	"DES_ENCRYPT()",
	"Dimension()",
	"Disjoint()",
	"Distance()",
	"DIV",
	"/",
	"ELT()",
	"ENCODE()",
	"ENCRYPT()",
	"EndPoint()",
	"Envelope()",
	"=",
	"<=>",
	"Equals()",
	"EXP()",
	"EXPORT_SET()",
	"ExteriorRing()",
	"EXTRACT()",
	"ExtractValue()",
	"FIELD()",
	"FIND_IN_SET()",
	"FIRST_VALUE()",
	"FLOOR()",
	"FORMAT()",
	"FORMAT_BYTES()",
	"FORMAT_PICO_TIME()",
	"FOUND_ROWS()",
	"FROM_BASE64()",
	"FROM_DAYS()",
	"FROM_UNIXTIME()",
	"GeomCollection()",
	"GeomCollFromText()",
	"GeometryCollectionFromText()",
	"GeomCollFromWKB()",
	"GeometryCollectionFromWKB()",
	"GeometryCollection()",
	"GeometryN()",
	"GeometryType()",
	"GeomFromText()",
	"GeometryFromText()",
	"GeomFromWKB()",
	"GeometryFromWKB()",
	"GET_DD_COLUMN_PRIVILEGES()",
	"GET_DD_CREATE_OPTIONS()",
	"GET_DD_INDEX_SUB_PART_LENGTH()",
	"GET_FORMAT()",
	"GET_LOCK()",
	"GLength()",
	">",
	">=",
	"GREATEST()",
	"GROUP_CONCAT()",
	"GROUPING()",
	"GTID_SUBSET()",
	"GTID_SUBTRACT()",
	"HEX()",
	"HOUR()",
	"ICU_VERSION()",
	"IF()",
	"IFNULL()",
	"IN()",
	"INET_ATON()",
	"INET_NTOA()",
	"INET6_ATON()",
	"INET6_NTOA()",
	"INSERT()",
	"INSTR()",
	"InteriorRingN()",
	"INTERNAL_AUTO_INCREMENT()",
	"INTERNAL_AVG_ROW_LENGTH()",
	"INTERNAL_CHECK_TIME()",
	"INTERNAL_CHECKSUM()",
	"INTERNAL_DATA_FREE()",
	"INTERNAL_DATA_LENGTH()",
	"INTERNAL_DD_CHAR_LENGTH()",
	"INTERNAL_GET_COMMENT_OR_ERROR()",
	"INTERNAL_GET_VIEW_WARNING_OR_ERROR()",
	"INTERNAL_INDEX_COLUMN_CARDINALITY()",
	"INTERNAL_INDEX_LENGTH()",
	"INTERNAL_KEYS_DISABLED()",
	"INTERNAL_MAX_DATA_LENGTH()",
	"INTERNAL_TABLE_ROWS()",
	"INTERNAL_UPDATE_TIME()",
	"Intersects()",
	"INTERVAL()",
	"IS",
	"IS_FREE_LOCK()",
	"IS_IPV4()",
	"IS_IPV4_COMPAT()",
	"IS_IPV4_MAPPED()",
	"IS_IPV6()",
	"IS NOT",
	"IS NOT NULL",
	"IS NULL",
	"IS_USED_LOCK()",
	"IS_UUID()",
	"IsClosed()",
	"IsEmpty()",
	"ISNULL()",
	"IsSimple()",
	"JSON_APPEND()",
	"JSON_ARRAY()",
	"JSON_ARRAY_APPEND()",
	"JSON_ARRAY_INSERT()",
	"JSON_ARRAYAGG()",
	"->",
	"JSON_CONTAINS()",
	"JSON_CONTAINS_PATH()",
	"JSON_DEPTH()",
	"JSON_EXTRACT()",
	"->>",
	"JSON_INSERT()",
	"JSON_KEYS()",
	"JSON_LENGTH()",
	"JSON_MERGE()",
	"JSON_MERGE_PATCH()",
	"JSON_MERGE_PRESERVE()",
	"JSON_OBJECT()",
	"JSON_OBJECTAGG()",
	"JSON_OVERLAPS()",
	"JSON_PRETTY()",
	"JSON_QUOTE()",
	"JSON_REMOVE()",
	"JSON_REPLACE()",
	"JSON_SCHEMA_VALID()",
	"JSON_SCHEMA_VALIDATION_REPORT()",
	"JSON_SEARCH()",
	"JSON_SET()",
	"JSON_STORAGE_FREE()",
	"JSON_STORAGE_SIZE()",
	"JSON_TABLE()",
	"JSON_TYPE()",
	"JSON_UNQUOTE()",
	"JSON_VALID()",
	"LAG()",
	"LAST_DAY",
	"LAST_INSERT_ID()",
	"LAST_VALUE()",
	"LCASE()",
	"LEAD()",
	"LEAST()",
	"LEFT()",
	"<<",
	"LENGTH()",
	"<",
	"<=",
	"LIKE",
	"LineFromText()",
	"LineStringFromText()",
	"LineFromWKB()",
	"LineStringFromWKB()",
	"LineString()",
	"LN()",
	"LOAD_FILE()",
	"LOCALTIME()",
	"LOCALTIME",
	"LOCALTIMESTAMP",
	"LOCALTIMESTAMP()",
	"LOCATE()",
	"LOG()",
	"LOG10()",
	"LOG2()",
	"LOWER()",
	"LPAD()",
	"LTRIM()",
	"MAKE_SET()",
	"MAKEDATE()",
	"MAKETIME()",
	"MASTER_POS_WAIT()",
	"MATCH",
	"MAX()",
	"MBRContains()",
	"MBRCoveredBy()",
	"MBRCovers()",
	"MBRDisjoint()",
	"MBREqual()",
	"MBREquals()",
	"MBRIntersects()",
	"MBROverlaps()",
	"MBRTouches()",
	"MBRWithin()",
	"MD",
	"MEMBER OF()",
	"MICROSECOND()",
	"MID()",
	"MIN()",
	"-",
	"MINUTE()",
	"MLineFromText()",
	"MultiLineStringFromText()",
	"MLineFromWKB()",
	"MultiLineStringFromWKB()",
	"MOD()",
	"%",
	"MOD",
	"MONTH()",
	"MONTHNAME()",
	"MPointFromText()",
	"MultiPointFromText()",
	"MPointFromWKB()",
	"MultiPointFromWKB()",
	"MPolyFromText()",
	"MultiPolygonFromText()",
	"MPolyFromWKB()",
	"MultiPolygonFromWKB()",
	"MultiLineString()",
	"MultiPoint()",
	"MultiPolygon()",
	"NAME_CONST()",
	"NOT",
	"!",
	"NOT BETWEEN ... AND ...",
	"!=",
	"<>",
	"NOT IN()",
	"NOT LIKE",
	"NOT REGEXP",
	"NOW()",
	"NTH_VALUE()",
	"NTILE()",
	"NULLIF()",
	"NumGeometries()",
	"NumInteriorRings()",
	"NumPoints()",
	"OCT()",
	"OCTET_LENGTH()",
	"OLD_PASSWORD()",
	"OR",
	"||",
	"ORD()",
	"Overlaps()",
	"PASSWORD()",
	"PERCENT_RANK()",
	"PERIOD_ADD()",
	"PERIOD_DIFF()",
	"PI()",
	"+",
	"Point()",
	"PointFromText()",
	"PointFromWKB()",
	"PointN()",
	"PolyFromText()",
	"PolygonFromText()",
	"PolyFromWKB()",
	"PolygonFromWKB()",
	"Polygon()",
	"POSITION()",
	"POW()",
	"POWER()",
	"PROCEDURE ANALYSE()",
	"PS_CURRENT_THREAD_ID()",
	"PS_THREAD_ID()",
	"QUARTER()",
	"QUOTE()",
	"RADIANS()",
	"RAND()",
	"RANDOM_BYTES()",
	"RANK()",
	"REGEXP",
	"REGEXP_INSTR()",
	"REGEXP_LIKE()",
	"REGEXP_REPLACE()",
	"REGEXP_SUBSTR()",
	"RELEASE_ALL_LOCKS()",
	"RELEASE_LOCK()",
	"REPEAT()",
	"REPLACE()",
	"REVERSE()",
	"RIGHT()",
	">>",
	"RLIKE",
	"ROLES_GRAPHML()",
	"ROUND()",
	"ROW_COUNT()",
	"ROW_NUMBER()",
	"RPAD()",
	"RTRIM()",
	"SCHEMA()",
	"SEC_TO_TIME()",
	"SECOND()",
	"SESSION_USER()",
	"SHA1()",
	"SHA()",
	"SHA2()",
	"SIGN()",
	"SIN()",
	"SLEEP()",
	"SOUNDEX()",
	"SOUNDS LIKE",
	"SPACE()",
	"SQL_THREAD_WAIT_AFTER_GTIDS()",
	"SQRT()",
	"SRID()",
	"ST_Area()",
	"ST_AsBinary()",
	"ST_AsWKB()",
	"ST_AsGeoJSON()",
	"ST_AsText()",
	"ST_AsWKT()",
	"ST_Buffer()",
	"ST_Buffer_Strategy()",
	"ST_Centroid()",
	"ST_Contains()",
	"ST_ConvexHull()",
	"ST_Crosses()",
	"ST_Difference()",
	"ST_Dimension()",
	"ST_Disjoint()",
	"ST_Distance()",
	"ST_Distance_Sphere()",
	"ST_EndPoint()",
	"ST_Envelope()",
	"ST_Equals()",
	"ST_ExteriorRing()",
	"ST_GeoHash()",
	"ST_GeomCollFromText()",
	"ST_GeometryCollectionFromText()",
	"ST_GeomCollFromTxt()",
	"ST_GeomCollFromWKB()",
	"ST_GeometryCollectionFromWKB()",
	"ST_GeometryN()",
	"ST_GeometryType()",
	"ST_GeomFromGeoJSON()",
	"ST_GeomFromText()",
	"ST_GeometryFromText()",
	"ST_GeomFromWKB()",
	"ST_GeometryFromWKB()",
	"ST_InteriorRingN()",
	"ST_Intersection()",
	"ST_Intersects()",
	"ST_IsClosed()",
	"ST_IsEmpty()",
	"ST_IsSimple()",
	"ST_IsValid()",
	"ST_LatFromGeoHash()",
	"ST_Latitude()",
	"ST_Length()",
	"ST_LineFromText()",
	"ST_LineStringFromText()",
	"ST_LineFromWKB()",
	"ST_LineStringFromWKB()",
	"ST_LongFromGeoHash()",
	"ST_Longitude()",
	"ST_MakeEnvelope()",
	"ST_MLineFromText()",
	"ST_MultiLineStringFromText()",
	"ST_MLineFromWKB()",
	"ST_MultiLineStringFromWKB()",
	"ST_MPointFromText()",
	"ST_MultiPointFromText()",
	"ST_MPointFromWKB()",
	"ST_MultiPointFromWKB()",
	"ST_MPolyFromText()",
	"ST_MultiPolygonFromText()",
	"ST_MPolyFromWKB()",
	"ST_MultiPolygonFromWKB()",
	"ST_NumGeometries()",
	"ST_NumInteriorRing()",
	"ST_NumInteriorRings()",
	"ST_NumPoints()",
	"ST_Overlaps()",
	"ST_PointFromGeoHash()",
	"ST_PointFromText()",
	"ST_PointFromWKB()",
	"ST_PointN()",
	"ST_PolyFromText()",
	"ST_PolygonFromText()",
	"ST_PolyFromWKB()",
	"ST_PolygonFromWKB()",
	"ST_Simplify()",
	"ST_SRID()",
	"ST_StartPoint()",
	"ST_SwapXY()",
	"ST_SymDifference()",
	"ST_Touches()",
	"ST_Transform()",
	"ST_Union()",
	"ST_Validate()",
	"ST_Within()",
	"ST_X()",
	"ST_Y()",
	"StartPoint()",
	"STATEMENT_DIGEST()",
	"STATEMENT_DIGEST_TEXT()",
	"STD()",
	"STDDEV()",
	"STDDEV_POP()",
	"STDDEV_SAMP()",
	"STR_TO_DATE()",
	"STRCMP()",
	"SUBDATE()",
	"SUBSTR()",
	"SUBSTRING()",
	"SUBSTRING_INDEX()",
	"SUBTIME()",
	"SUM()",
	"SYSDATE()",
	"SYSTEM_USER()",
	"TAN()",
	"TIME()",
	"TIME_FORMAT()",
	"TIME_TO_SEC()",
	"TIMEDIFF()",
	"*",
	"TIMESTAMP()",
	"TIMESTAMPADD()",
	"TIMESTAMPDIFF()",
	"TO_BASE64()",
	"TO_DAYS()",
	"TO_SECONDS()",
	"Touches()",
	"TRIM()",
	"TRUNCATE()",
	"UCASE()",
	"-",
	"UNCOMPRESS()",
	"UNCOMPRESSED_LENGTH()",
	"UNHEX()",
	"UNIX_TIMESTAMP()",
	"UpdateXML()",
	"UPPER()",
	"USER()",
	"UTC_DATE()",
	"UTC_TIME()",
	"UTC_TIMESTAMP()",
	"UUID()",
	"UUID_SHORT()",
	"UUID_TO_BIN()",
	"VALIDATE_PASSWORD_STRENGTH()",
	"VALUES()",
	"VAR_POP()",
	"VAR_SAMP()",
	"VARIANCE()",
	"VERSION()",
	"WAIT_FOR_EXECUTED_GTID_SET()",
	"WAIT_UNTIL_SQL_THREAD_AFTER_GTIDS()",
	"WEEK()",
	"WEEKDAY()",
	"WEEKOFYEAR()",
	"WEIGHT_STRING()",
	"Within()",
	"X()",
	"XOR",
	"Y()",
	"YEAR()",
	"YEARWEEK()",
}