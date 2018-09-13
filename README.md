# modbus-service
Go binding for libmodbus with GRPC server acting as a modbus master running in docker
## Linux Shared library
### libmodbus 3.1.4
Currently libmodbus 3.1.4 is built and installed in a linux container with a debian image.
Debian was chosen over alpine because it depends on serial.h for the RTU(Serial) implementation.  It gets installed in /usr/local/lib and 
and the ENV LD_LIBRARY_PATH is set to tell the go compiler where the lib is located.

## GRPC API
Currently only TCP is supported

#### ReadHoldingReg
Reads a set of Registers 
#### ReadCoil
Reads a set of Coils
#### OpenTcpConnection
Opens a single tcp connection 
#### CloseConnection
Closes that connection
#### GetConnectionStatus
Gets connection info