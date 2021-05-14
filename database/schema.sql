USE master
GO

IF NOT EXISTS (
    SELECT [name]
        FROM [sys].[databases]
        WHERE name = N'lms'
)
CREATE DATABASE [lms]
GO

USE [lms]
GO


