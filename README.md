Multi -Backend Caching Library in Go

1.	Overview:
    
This project aims to develop a robust caching library in Go that supports multiple backends. The library will feature an in-memory cache with an LRU (Least Recently Used) eviction policy and integrate with external caching solutions (Redis). It will provide an easy-to-use API for cache operations and include cache invalidation and expiration policies.
The library aims to offer a flexible and performant caching solution include LRU and Redis cache storage options, helping developers improve application performance and scalability by reducing latency and external database loads.

Key Features:

⦁	In-memory LRU Cache for quick data retrieval with O(1) time complexity for get/set operations.

⦁	Redis Integration for distributed and persistent caching.

⦁	Unified API for easy integration, with support for cache expiration (TTL) and invalidation.

⦁	High Performance and Scalability due to efficient memory usage and external cache offloading.

2. API Documentation:
   
Overview: 

This API allows users to perform caching operations using both an in-memory LRU cache and an external Redis cache. It supports setting, getting, and deleting cache entries, along with clearing all cache data.

Base URL:
http://localhost:8080

Endpoints:

1. Set Cache entry:
   
⦁	Endpoint: /set

⦁	Method: POST

⦁	Description: Sets a key-value pair in both the LRU cache and Redis cache with TTL (Time-To-Live).

⦁	Request Body: 

           {
            "key": "string",          
            "value": "interface{}",    
            "ttl": "int"               
            }

2. Get Cache entry:
   
⦁	Endpoint: /get/:key

⦁	Method: GET

⦁	Description: Retrieves a value from the cache by key.

⦁	Response:

          {
            "value": "cached value",
            "source": "LRU"
          }
          
3. Get All Cache entry:
   
⦁	Endpoint: /get

⦁	Method: GET

⦁	Description: Retrieves all values from the cache.

Response:

          {
            "value": "cached value",
            "source": "LRU"
          }
          
4. Delete Cache entry:
   
⦁	Endpoint: /delete/:key

⦁	Method: DELETE

⦁	Description: Deletes a cache entry from both the LRU cache and Redis cache.

⦁	Response:

          {
           "status": "delete successful"
          }

5. Clear Cache:
   
⦁	Endpoint: /clear

⦁	Method: POST

⦁	Description: Clears all cache entries from both the LRU cache and Redis cache.

⦁	Response:

          {
            "status": "cache cleared"
          }
          
Error Codes:

⦁	400 Bad Request: The client has sent a request that cannot be processed.

⦁	404 Not Found: The requested key does not exist in the cache.

⦁	500 Internal Server Error: There was an internal error in processing the request (e.g., failure to interact with Redis).

3. Usage Guide:
   
Overview:

This API provides endpoints for interacting with a caching system that uses both an in-memory LRU cache and an external Redis cache. We can perform operations such as setting, getting, deleting, and clearing cache entries.

Endpoints and Examples

Postman Example:

1. Post cache entry:
   
⦁	Open Postman and create a new request.

⦁	Set the request method to POST.

⦁	Enter the URL http://localhost:8080/set.

⦁	Go to the "Headers" tab and set Content-Type to application/json.

⦁	Go to the "Body" tab, select raw, and choose JSON format.

⦁	Enter the JSON body:

         {
          "key": "exampleKey",
          "value": "exampleValue",
          "ttl": 60
         }
         
⦁	Click "Send" to execute the request.

⦁	Response:

200 OK:

         {
          "status": "set successful"
          }
          
400 Bad Request:

        {   
         "error": "error message"
         } 
         
500 Internal Server Error:

        {
          "error": "error message"
         }
         
2. Get Cache Entry:
   
Postman Example:

⦁	Open Postman and create a new request.

⦁	Set the request method to GET.

⦁	Enter the URL http://localhost:8080/get/Key.

⦁	Click "Send" to execute the request.

⦁	Response:

200 OK (if found):

       {
        "value": "exampleValue",
        "source": "LRU" or "Redis" depending on where it was found
       }
404 Not Found:

       {
        "error": "key not found"
       }
       
3. Get All Cache Entry:
   
Postman Example:

⦁	Open Postman and create a new request.

⦁	Set the request method to GET.

⦁	Enter the URL http://localhost:8080/get.

⦁	Click "Send" to execute the request.

⦁	Response:

200 OK (if found):

        {
         "value": "exampleValue",
         "source": "LRU" or "Redis" depending on where it was found
        }
        
404 Not Found:

        {
         "error": "key not found"
        }
        
4. Delete Cache Entry:
   
Postman Example:

⦁	Open Postman and create a new request.

⦁	Set the request method to DELETE.

⦁	Enter the URL http://localhost:8080/delete/Key.

⦁	Click "Send" to execute the request.

⦁	Response:

200 OK:

       {
        "status": "delete successful"
       }
       
404 Not Found:

       {
        "error": "error message"
       }
       
500 Internal Server Error:

       {
        "error": "error message"
       }
       
5. Clear Cache:
    
Postman Example:

⦁	Open Postman and create a new request.

⦁	Set the request method to POST.

⦁	Enter the URL http://localhost:8080/clear.

⦁	Click "Send" to execute the request.

⦁	Response:

200 OK:

      {
       "status": "cache cleared"
      }
      
500 Internal Server Error:

      {
       "error": "error message"
      }
      
4. Unit Testing, Integration Testing and Benchmark Results:
   
LRU Cache Testing:

⦁	Set and Get Functionality:

Functionality: The LRU cache correctly sets and retrieves values.

Tests: The SetAndGet test confirmed that values can be stored and retrieved as expected. Expired entries are correctly handled, and eviction works as intended when the cache exceeds its capacity.

⦁	Expiration Handling:

Functionality: The cache properly handles expired entries.

Tests: The cache accurately reports a cache miss for entries that have expired, as observed in the TestLRU_SetAndGet test.

⦁	Eviction:

Functionality: The LRU cache evicts the least recently used entries when capacity is exceeded.

Tests: The TestLRU_SetAndGet test demonstrated that the oldest entry (based on usage) is removed when a new entry is added and the cache is full.

⦁	Deletion:

Functionality: The cache successfully deletes entries.

Tests: The TestLRU_Delete test showed that entries can be removed and are no longer retrievable.

⦁	Clearing:

Functionality: The Clear function effectively removes all entries from the cache.

Tests: The TestLRU_Clear test confirmed that the cache can be cleared, and all entries are removed.

⦁	GetAllKeys Functionality:

Functionality: The GetAllKeys function retrieves all current keys and values in the cache.

Tests: The TestLRU_GetAllKeys test showed that all keys are returned and matched the expected keys.

Redis Cache Testing:

⦁	Set and Get Functionality:

Functionality: The Redis cache correctly sets and retrieves JSON-encoded values.

Tests: The TestRedisCache set and get tests confirmed that data is correctly stored and retrieved, and data integrity is maintained.

⦁	Deletion:

Functionality: Redis cache supports entry deletion.

Tests: The TestRedisCache_Delete test confirmed that entries can be deleted and are no longer retrievable.

⦁	Clearing:

Functionality: The Clear function in Redis effectively removes all entries.

Tests: The TestRedisCache_Clear test showed that all cache entries are cleared when the cache is flushed.

⦁	GetAllKeys Functionality:

Functionality: Redis supports retrieving all keys and values.

Tests: The TestRedisCache_GetAllKeys test demonstrated that all keys in the cache can be retrieved, matching the expected results.

<img width="382" alt="testing" src="https://github.com/user-attachments/assets/2ba0c37f-a49b-47ca-8ad4-e48ad3da76c7">

LRU Cache Benchmarking:

⦁	Set and Get Performance:

Benchmarks reveal the performance for setting and retrieving entries in the LRU cache. Since it is fully in-memory and operates locally, the Set and Get operations are very fast, with minimal overhead from managing the eviction list.

⦁	GetAllKeys Performance:

Benchmarks for GetAllKeys measure how efficiently the LRU cache handles retrieving all keys and values stored in the cache. The performance is dependent on the internal data structure and the cache size, though it is generally quick as everything is in-memory.

⦁	Delete and Clear Performance:

Benchmarks for Delete reflect the speed of removing a specific key from the LRU cache, involving both the removal from the list and map. Clear benchmarks show how fast the LRU cache can purge all entries, which is generally efficient given that it’s an in-memory operation.

Redis Cache Benchmarking:

⦁	Set and Get Performance:

Benchmark tests show the performance of setting and getting individual entries in Redis. Since Redis is an in-memory key-value store, these operations are generally efficient, but they may involve slight overhead due to network communication and serialization.

⦁	GetAllKeys Performance:

Benchmark tests for GetAllKeys provide insights into the performance of retrieving all keys and values from Redis. This operation can be impacted by the number of keys and the server's performance, though Redis is typically optimized for handling such queries efficiently.

⦁	Delete and Clear Performance:

Benchmarks for Delete show how fast Redis can remove individual keys, which typically involves a straightforward key removal process. Clear benchmarks test Redis's ability to flush the entire database, which is optimized but can be slower depending on the number of stored entries.

<img width="467" alt="benchmarking" src="https://github.com/user-attachments/assets/fdc40888-548e-454d-a88c-e75e85a319b6">

5. Best practices for integrating the library into Go applications:
   
⦁	Initialize Caches Properly

Initialization: Ensure that both the in-memory LRU cache and the Redis cache are initialized properly in our application's setup or initialization code. Use a singleton pattern or dependency injection to manage instances.

⦁	Use Cache Interfaces

Abstraction: Define common interfaces for cache operations (e.g., Set, Get, Delete) to abstract away the underlying implementation. This allows us to switch between different cache backends with minimal code changes.

⦁	Implement Error Handling

Robust Error Handling: Ensure that our code handles errors gracefully. When interacting with caches, handle errors such as connectivity issues or timeouts and provide meaningful feedback or fallback mechanisms.

⦁	Optimize Performance

TTL Management: Configure appropriate TTL values for cache entries to balance between memory usage and performance. Regularly monitor and adjust TTL settings based on application usage patterns.

Cache Eviction: Ensure that the LRU cache's maximum size is configured based on the expected load and available memory to avoid excessive memory usage.

⦁	Monitor and Log

Monitoring: Implement monitoring for cache hits, misses, and errors. 

Logging: Include logging for cache operations to aid in debugging and performance tuning. Log cache operations at an appropriate level to avoid excessive log volume.

⦁	Testing

Unit Testing: Write unit tests for each cache implementation to ensure correct behavior. Mock external dependencies (like Redis) to isolate and test the in-memory cache functionality.

Integration Testing: Test interactions between our application and the caching library to ensure that both LRU and Redis caches work as expected in a real-world scenario.

Benchmarking: Utilize Go’s built-in benchmarking framework (testing package) to write benchmark tests. This allows us to measure the performance of cache operations under different scenarios.

⦁	Documentation

Documentation: Provide comprehensive documentation for the caching library, including usage examples, configuration options, and best practices. Ensure that developers understand how to integrate and use the library effectively.
