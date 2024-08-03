package utils

const DbNameRedis = "redis"
const DbNameMongoDB = "mongodb"
const DbNameCassandra = "cassandra"
const DbNameMySQL = "mysql"

const DenormalizationFactor = 16
const WorkerCount = 250

type TestType string

const TestTypeRead TestType = "readheavy"
const TestTypeBalanced TestType = "balanced"
const TestTypeWrite TestType = "writeheavy"

type TestDataType string

const TestDataTypeSm TestDataType = "sm"
const TestDataTypeLg TestDataType = "lg"
const TestDataTypeImg TestDataType = "img"

type OpType string

const OpTypeGet OpType = "GET"
const OpTypePut OpType = "PUT"

const WriteFactorRead = 0.1
const WriteFactorBalanced = 0.5
const WriteFactorWrite = 0.9
