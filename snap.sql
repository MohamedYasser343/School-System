-- phpMyAdmin SQL Dump
-- version 4.9.5deb2
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Mar 29, 2022 at 09:25 AM
-- Server version: 8.0.28-0ubuntu0.20.04.3
-- PHP Version: 7.4.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `snap`
--

-- --------------------------------------------------------

--
-- Table structure for table `absence`
--

CREATE TABLE `absence` (
  `absenceId` int NOT NULL,
  `stdId` int DEFAULT NULL,
  `morningLine` tinyint(1) DEFAULT NULL,
  `session1` tinyint(1) DEFAULT NULL,
  `session2` tinyint(1) DEFAULT NULL,
  `session3` tinyint(1) DEFAULT NULL,
  `session4` tinyint(1) DEFAULT NULL,
  `session5` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `academic_syllabus`
--

CREATE TABLE `academic_syllabus` (
  `academic_syllabus_id` int NOT NULL,
  `academic_syllabus_code` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `title` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `description` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `class_id` int DEFAULT NULL,
  `uploader_type` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `uploader_id` int DEFAULT NULL,
  `year` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `timestamp` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `file_name` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `subject_id` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `accountant`
--

CREATE TABLE `accountant` (
  `accountant_id` int NOT NULL,
  `name` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `email` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `password` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `addteaviolation`
--

CREATE TABLE `addteaviolation` (
  `addteaviolationid` int NOT NULL,
  `violationname` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `addvacation`
--

CREATE TABLE `addvacation` (
  `addvacationid` int NOT NULL,
  `vacationname` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `addvacation`
--
-- -----------------------------------------------------------
--
-- Table structure for table `admin`
--

CREATE TABLE `admin` (
  `admin_id` int NOT NULL,
  `name` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `email` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `password` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `level` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `authentication_key` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `phone` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `address` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `AssignScholar`
--

CREATE TABLE `AssignScholar` (
  `AssignScholarId` int NOT NULL,
  `StdId` int NOT NULL,
  `ScholarshipId` int NOT NULL,
  `ScholarShipKind` varchar(300) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--

-- --------------------------------------------------------

--
-- Table structure for table `attendance`
--

CREATE TABLE `attendance` (
  `attendance_id` int NOT NULL,
  `timestamp` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `year` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `class_id` int DEFAULT NULL,
  `section_id` int DEFAULT NULL,
  `student_id` int DEFAULT NULL,
  `class_routine_id` int DEFAULT NULL,
  `status` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `book`
--

CREATE TABLE `book` (
  `book_id` int NOT NULL,
  `name` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `description` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `author` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `class_id` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `price` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `total_copies` int DEFAULT NULL,
  `issued_copies` int DEFAULT NULL,
  `status` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `file_name` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `book_request`
--

CREATE TABLE `book_request` (
  `book_request_id` int NOT NULL,
  `book_id` int DEFAULT NULL,
  `student_id` int DEFAULT NULL,
  `issue_start_date` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `issue_end_date` longtext CHARACTER SET utf8 COLLATE utf8_unicode_ci,
  `status` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `bus`
--

CREATE TABLE `bus` (
  `busid` int NOT NULL,
  `busline` varchar(200) DEFAULT NULL,
  `drivername` varchar(200) DEFAULT NULL,
  `driverNo` varchar(11) DEFAULT NULL,
  `startpoint` varchar(200) DEFAULT NULL,
  `starttime` varchar(200) DEFAULT NULL,
  `busChairno` int DEFAULT NULL,
  `endpoint` varchar(200) DEFAULT NULL,
  `endtime` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--

-- --------------------------------------------------------

--
-- Table structure for table `bussing`
--

CREATE TABLE `bussing` (
  `bussingid` int NOT NULL,
  `busid` int NOT NULL,
  `stdid` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `capstone`
--

CREATE TABLE `capstone` (
  `capstoneid` int NOT NULL,
  `groupid` varchar(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--

--

-- --------------------------------------------------------

--
-- Table structure for table `ci_sessions`
--

CREATE TABLE `ci_sessions` (
  `id` varchar(40) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `ip_address` varchar(45) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
  `timestamp` int UNSIGNED NOT NULL DEFAULT '0',
  `data` blob NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COLLATE=utf8_unicode_ci;

-- --------------------------------------------------------

--
-- Table structure for table `Class`
--

CREATE TABLE `Class` (
  `ClassId` int NOT NULL,
  `ClassName` varchar(200) DEFAULT NULL,
  `ClassGrade` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--

-- --------------------------------------------------------

--
-- Table structure for table `ClassStd`
--

CREATE TABLE `ClassStd` (
  `ClassStdId` int NOT NULL,
  `ClassId` int NOT NULL,
  `StdId` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `competition`
--

CREATE TABLE `competition` (
  `competitionid` int NOT NULL,
  `name` varchar(200) DEFAULT NULL,
  `sponser` varchar(200) DEFAULT NULL,
  `date` date DEFAULT NULL,
  `notes` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `entercompetition`
--

CREATE TABLE `entercompetition` (
  `entercompid` int NOT NULL,
  `competitionid` int NOT NULL,
  `stdid` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `overlooking`
--

CREATE TABLE `overlooking` (
  `overlookingid` int NOT NULL,
  `overlookingname` varchar(200) DEFAULT NULL,
  `overlookingmobile` varchar(11) DEFAULT NULL,
  `overlookingnationalid` varchar(14) DEFAULT NULL,
  `overlookingaddr` varchar(200) DEFAULT NULL,
  `educationalqualification` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--

-- --------------------------------------------------------

--
-- Table structure for table `parent`
--

CREATE TABLE `parent` (
  `parentId` int NOT NULL,
  `studentId` int NOT NULL,
  `parentName` varchar(200) NOT NULL,
  `parentNationalId` varchar(200) NOT NULL,
  `parentJob` varchar(200) NOT NULL,
  `parentMobile` varchar(200) NOT NULL,
  `parentMobile2` varchar(200) NOT NULL,
  `parentAddr` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `permission`
--

CREATE TABLE `permission` (
  `permissionid` int NOT NULL,
  `permissionname` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `recordstdcapstone`
--

CREATE TABLE `recordstdcapstone` (
  `recordstdcapstoneid` int NOT NULL,
  `capstoneid` int NOT NULL,
  `studentid` int NOT NULL,
  `studentid2` int NOT NULL,
  `studentid3` int NOT NULL,
  `studentid4` int NOT NULL,
  `studentid5` int NOT NULL,
  `studentid6` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `recordstduni`
--

CREATE TABLE `recordstduni` (
  `recordstduniid` int NOT NULL,
  `stdid` int DEFAULT NULL,
  `uniid` int DEFAULT NULL,
  `unifees` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `recordstdviolation`
--

CREATE TABLE `recordstdviolation` (
  `recordstdviolationid` int NOT NULL,
  `studentid` int NOT NULL,
  `violationid` int NOT NULL,
  `violationdate` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--

-- --------------------------------------------------------

--
-- Table structure for table `recordteaviolation`
--

CREATE TABLE `recordteaviolation` (
  `recordteaviolationid` int NOT NULL,
  `teacherid` int NOT NULL,
  `teaviolationname` int DEFAULT NULL,
  `teaviolationdate` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `recordvacation`
--

CREATE TABLE `recordvacation` (
  `recordvacationid` int NOT NULL,
  `teacherid` int NOT NULL,
  `vacationtype` int NOT NULL,
  `vacationfrom` varchar(200) DEFAULT NULL,
  `vacationto` varchar(200) DEFAULT NULL,
  `numdays` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `romming`
--

CREATE TABLE `romming` (
  `rommingid` int NOT NULL,
  `roomid` int NOT NULL,
  `stdid` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `room`
--

CREATE TABLE `room` (
  `roomid` int NOT NULL,
  `roomno` varchar(5) DEFAULT NULL,
  `maxnO` int DEFAULT NULL,
  `floorno` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `scholarship`
--

CREATE TABLE `scholarship` (
  `scolarshipid` int NOT NULL,
  `universityname` varchar(200) DEFAULT NULL,
  `collagename` varchar(200) DEFAULT NULL,
  `notes` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--

-- --------------------------------------------------------

--
-- Table structure for table `school`
--

CREATE TABLE `school` (
  `schoolid` int NOT NULL,
  `schoolname` varchar(200) DEFAULT NULL,
  `year` date DEFAULT NULL,
  `princname` varchar(200) DEFAULT NULL,
  `depname` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `socialsharing`
--

CREATE TABLE `socialsharing` (
  `socialsharingid` int NOT NULL,
  `parentid` int DEFAULT NULL,
  `sharingtype` varchar(200) DEFAULT NULL,
  `sharingnum` int DEFAULT NULL,
  `cash` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `stem`
--

CREATE TABLE `stem` (
  `stemid` int NOT NULL,
  `studentid` int DEFAULT NULL,
  `stdorder` int DEFAULT NULL,
  `stdstemid` int DEFAULT NULL,
  `lapinsur` int DEFAULT NULL,
  `lapvouch` int DEFAULT NULL,
  `lapdate` date DEFAULT NULL,
  `stdfees` int DEFAULT NULL,
  `stddate` date DEFAULT NULL,
  `stdclass` int NOT NULL,
  `secondlang` varchar(200) DEFAULT NULL,
  `activity` varchar(200) DEFAULT NULL,
  `groupid` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `student`
--

CREATE TABLE `student` (
  `StdId` int NOT NULL,
  `Code` int DEFAULT NULL,
  `English Name` varchar(49) DEFAULT NULL,
  `Arabic Name` varchar(41) DEFAULT NULL,
  `Grade` varchar(3) DEFAULT NULL,
  `specilization` varchar(10) DEFAULT NULL,
  `SecondLang` varchar(6) DEFAULT NULL,
  `Religion` varchar(9) DEFAULT NULL,
  `Gender` varchar(4) DEFAULT NULL,
  `Portfolio No` int DEFAULT NULL,
  `National ID (14 digit)` bigint DEFAULT NULL,
  `Class` varchar(2) DEFAULT NULL,
  `PersonalEmail` varchar(35) DEFAULT NULL,
  `Email` varchar(42) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--

-- --------------------------------------------------------

--
-- Table structure for table `StudentViolations`
--

CREATE TABLE `StudentViolations` (
  `Name` varchar(200) NOT NULL,
  `Id` varchar(200) NOT NULL,
  `Violation1` varchar(200) NOT NULL,
  `Violation2` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--

-- --------------------------------------------------------

--
-- Table structure for table `student_old`
--

CREATE TABLE `student_old` (
  `studentid2` int NOT NULL,
  `stdname` varchar(200) DEFAULT NULL,
  `stdnamee` varchar(200) DEFAULT NULL,
  `stdpermesion` int DEFAULT NULL,
  `stdnationalid` varchar(200) DEFAULT NULL,
  `stdgender` varchar(200) DEFAULT NULL,
  `stdreligon` varchar(200) DEFAULT NULL,
  `stdmail` varchar(200) DEFAULT NULL,
  `stdstate` varchar(200) DEFAULT NULL,
  `stdaddr` varchar(200) DEFAULT NULL,
  `stdschool` varchar(200) DEFAULT NULL,
  `stdschoolType` varchar(200) DEFAULT NULL,
  `stdschoolProvince` varchar(200) DEFAULT NULL,
  `stdschoolCity` varchar(200) DEFAULT NULL,
  `stdmobile` varchar(200) DEFAULT NULL,
  `stdmobile2` varchar(200) DEFAULT NULL,
  `DadDeth` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `dadname` varchar(200) DEFAULT NULL,
  `dadnationalid` varchar(200) DEFAULT NULL,
  `dadjob` varchar(200) DEFAULT NULL,
  `dadmobile` varchar(200) DEFAULT NULL,
  `dadmobile2` varchar(200) DEFAULT NULL,
  `dadaddr` varchar(200) DEFAULT NULL,
  `MumDeth` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `mumname` varchar(200) DEFAULT NULL,
  `mumnationalid` varchar(200) DEFAULT NULL,
  `mumjob` varchar(200) DEFAULT NULL,
  `mummobile` varchar(200) DEFAULT NULL,
  `mummobile2` varchar(200) DEFAULT NULL,
  `mumaddr` varchar(200) DEFAULT NULL,
  `stdbirthcity` varchar(200) DEFAULT NULL,
  `stdbirthyear` varchar(200) DEFAULT NULL,
  `stdday` varchar(200) DEFAULT NULL,
  `stdmonth` varchar(200) DEFAULT NULL,
  `stdyear` varchar(200) DEFAULT NULL,
  `stdphotolink` varchar(200) DEFAULT NULL,
  `UserName` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `Password` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `teachers`
--

CREATE TABLE `teachers` (
  `teachersId` int NOT NULL,
  `teacherCode` int DEFAULT NULL,
  `TeacherName` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `JobCadr` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `SupervisingPosOriginSchool` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `SupervisingPosStemSchool` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `EduLevel` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `HiringDate` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `OriginQualification` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `ObtainedOriginQualification` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `QualificationDate` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `Specialization` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `OriginQualificationGrade` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `HighQualification` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `ObtainedHighQualification` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `HighQualificationDate` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `HighQualificationGrade` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `FinancialGrade` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `FinancialGradeDate` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `OriginSchool` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `OriginTeachingGov` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `OriginCity` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `StemSchoolName` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `CurrentStemSchoolEnterDate` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `StemSchoolEnterDate` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `MobileWhats` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `MobileCalls` varchar(200) DEFAULT NULL,
  `Email` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `NationalId` varchar(200) DEFAULT NULL,
  `Photo` varchar(200) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

--

-- --------------------------------------------------------

--
-- Table structure for table `trip`
--

CREATE TABLE `trip` (
  `tripid` int NOT NULL,
  `tripname` varchar(200) DEFAULT NULL,
  `tripdate` varchar(200) DEFAULT NULL,
  `tripplace` varchar(200) DEFAULT NULL,
  `tripchairnum` int DEFAULT NULL,
  `triptype` varchar(200) DEFAULT NULL,
  `publicoverlooking` varchar(200) DEFAULT NULL,
  `overlooking` varchar(200) DEFAULT NULL,
  `overlooking2` varchar(200) DEFAULT NULL,
  `overlooking3` varchar(200) DEFAULT NULL,
  `overlooking4` varchar(200) DEFAULT NULL,
  `overlooking5` varchar(200) DEFAULT NULL,
  `takeofftime` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--

-- --------------------------------------------------------

--
-- Table structure for table `University`
--

CREATE TABLE `University` (
  `UniversityID` int NOT NULL,
  `UniversityName` varchar(200) DEFAULT NULL,
  `CollageName` varchar(200) DEFAULT NULL,
  `CIty` varchar(200) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint UNSIGNED NOT NULL,
  `name` longtext,
  `email` longtext,
  `password` longtext,
  `permission` bigint DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--

-- --------------------------------------------------------

--
-- Table structure for table `violations`
--

CREATE TABLE `violations` (
  `violationid` int NOT NULL,
  `violationName` varchar(200) DEFAULT NULL,
  `violationPoint` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--

--
-- Indexes for table `absence`
--
ALTER TABLE `absence`
  ADD PRIMARY KEY (`absenceId`),
  ADD KEY `stdId` (`stdId`);

--
-- Indexes for table `academic_syllabus`
--
ALTER TABLE `academic_syllabus`
  ADD PRIMARY KEY (`academic_syllabus_id`);

--
-- Indexes for table `accountant`
--
ALTER TABLE `accountant`
  ADD PRIMARY KEY (`accountant_id`);

--
-- Indexes for table `addteaviolation`
--
ALTER TABLE `addteaviolation`
  ADD PRIMARY KEY (`addteaviolationid`);

--
-- Indexes for table `addvacation`
--
ALTER TABLE `addvacation`
  ADD PRIMARY KEY (`addvacationid`);

--
-- Indexes for table `admin`
--
ALTER TABLE `admin`
  ADD PRIMARY KEY (`admin_id`);

--
-- Indexes for table `AssignScholar`
--
ALTER TABLE `AssignScholar`
  ADD PRIMARY KEY (`AssignScholarId`),
  ADD KEY `StdId` (`StdId`),
  ADD KEY `ScholarshipId` (`ScholarshipId`);

--
-- Indexes for table `attendance`
--
ALTER TABLE `attendance`
  ADD PRIMARY KEY (`attendance_id`);

--
-- Indexes for table `book`
--
ALTER TABLE `book`
  ADD PRIMARY KEY (`book_id`);

--
-- Indexes for table `book_request`
--
ALTER TABLE `book_request`
  ADD PRIMARY KEY (`book_request_id`);

--
-- Indexes for table `bus`
--
ALTER TABLE `bus`
  ADD PRIMARY KEY (`busid`);

--
-- Indexes for table `bussing`
--
ALTER TABLE `bussing`
  ADD PRIMARY KEY (`bussingid`),
  ADD KEY `busid` (`busid`) USING BTREE,
  ADD KEY `stdid` (`stdid`) USING BTREE;

--
-- Indexes for table `capstone`
--
ALTER TABLE `capstone`
  ADD PRIMARY KEY (`capstoneid`);

--
-- Indexes for table `ci_sessions`
--
ALTER TABLE `ci_sessions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `ci_sessions_timestamp` (`timestamp`);

--
-- Indexes for table `Class`
--
ALTER TABLE `Class`
  ADD PRIMARY KEY (`ClassId`);

--
-- Indexes for table `ClassStd`
--
ALTER TABLE `ClassStd`
  ADD PRIMARY KEY (`ClassStdId`),
  ADD KEY `ClassId` (`ClassId`),
  ADD KEY `StdId` (`StdId`);

--
-- Indexes for table `competition`
--
ALTER TABLE `competition`
  ADD PRIMARY KEY (`competitionid`);

--
-- Indexes for table `entercompetition`
--
ALTER TABLE `entercompetition`
  ADD PRIMARY KEY (`entercompid`),
  ADD KEY `competitionid` (`competitionid`),
  ADD KEY `stdid` (`stdid`);

--
-- Indexes for table `overlooking`
--
ALTER TABLE `overlooking`
  ADD PRIMARY KEY (`overlookingid`);

--
-- Indexes for table `parent`
--
ALTER TABLE `parent`
  ADD PRIMARY KEY (`parentId`),
  ADD KEY `studentId` (`studentId`);

--
-- Indexes for table `permission`
--
ALTER TABLE `permission`
  ADD PRIMARY KEY (`permissionid`);

--
-- Indexes for table `recordstdcapstone`
--
ALTER TABLE `recordstdcapstone`
  ADD PRIMARY KEY (`recordstdcapstoneid`),
  ADD KEY `capstoneid` (`capstoneid`),
  ADD KEY `studentid` (`studentid`),
  ADD KEY `studentid2` (`studentid2`),
  ADD KEY `studentid3` (`studentid3`),
  ADD KEY `studentid4` (`studentid4`),
  ADD KEY `studentid5` (`studentid5`),
  ADD KEY `studentid6` (`studentid6`);

--
-- Indexes for table `recordstduni`
--
ALTER TABLE `recordstduni`
  ADD PRIMARY KEY (`recordstduniid`),
  ADD KEY `stdid` (`stdid`),
  ADD KEY `uniid` (`uniid`);

--
-- Indexes for table `recordstdviolation`
--
ALTER TABLE `recordstdviolation`
  ADD PRIMARY KEY (`recordstdviolationid`),
  ADD KEY `studentid` (`studentid`),
  ADD KEY `violationid` (`violationid`);

--
-- Indexes for table `recordteaviolation`
--
ALTER TABLE `recordteaviolation`
  ADD PRIMARY KEY (`recordteaviolationid`),
  ADD KEY `teacherid` (`teacherid`),
  ADD KEY `teaviolationname` (`teaviolationname`);

--
-- Indexes for table `recordvacation`
--
ALTER TABLE `recordvacation`
  ADD PRIMARY KEY (`recordvacationid`),
  ADD KEY `teacherid` (`teacherid`),
  ADD KEY `vacationtype` (`vacationtype`);

--
-- Indexes for table `romming`
--
ALTER TABLE `romming`
  ADD PRIMARY KEY (`rommingid`),
  ADD KEY `roomid` (`roomid`),
  ADD KEY `stdid` (`stdid`);

--
-- Indexes for table `room`
--
ALTER TABLE `room`
  ADD PRIMARY KEY (`roomid`);

--
-- Indexes for table `scholarship`
--
ALTER TABLE `scholarship`
  ADD PRIMARY KEY (`scolarshipid`);

--
-- Indexes for table `school`
--
ALTER TABLE `school`
  ADD PRIMARY KEY (`schoolid`);

--
-- Indexes for table `socialsharing`
--
ALTER TABLE `socialsharing`
  ADD PRIMARY KEY (`socialsharingid`),
  ADD KEY `parentid` (`parentid`);

--
-- Indexes for table `stem`
--
ALTER TABLE `stem`
  ADD PRIMARY KEY (`stemid`),
  ADD KEY `stdclass` (`stdclass`),
  ADD KEY `groupid` (`groupid`),
  ADD KEY `studentid` (`studentid`);

--
-- Indexes for table `student`
--
ALTER TABLE `student`
  ADD PRIMARY KEY (`StdId`);

--
-- Indexes for table `student_old`
--
ALTER TABLE `student_old`
  ADD PRIMARY KEY (`studentid2`),
  ADD KEY `stdpermesion` (`stdpermesion`);

--
-- Indexes for table `teachers`
--
ALTER TABLE `teachers`
  ADD PRIMARY KEY (`teachersId`);

--
-- Indexes for table `trip`
--
ALTER TABLE `trip`
  ADD PRIMARY KEY (`tripid`);

--
-- Indexes for table `University`
--
ALTER TABLE `University`
  ADD PRIMARY KEY (`UniversityID`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `violations`
--
ALTER TABLE `violations`
  ADD PRIMARY KEY (`violationid`),
  ADD KEY `violationName` (`violationName`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `absence`
--
ALTER TABLE `absence`
  MODIFY `absenceId` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `academic_syllabus`
--
ALTER TABLE `academic_syllabus`
  MODIFY `academic_syllabus_id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `accountant`
--
ALTER TABLE `accountant`
  MODIFY `accountant_id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `addteaviolation`
--
ALTER TABLE `addteaviolation`
  MODIFY `addteaviolationid` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `addvacation`
--
ALTER TABLE `addvacation`
  MODIFY `addvacationid` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `admin`
--
ALTER TABLE `admin`
  MODIFY `admin_id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `AssignScholar`
--
ALTER TABLE `AssignScholar`
  MODIFY `AssignScholarId` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `attendance`
--
ALTER TABLE `attendance`
  MODIFY `attendance_id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `book`
--
ALTER TABLE `book`
  MODIFY `book_id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `book_request`
--
ALTER TABLE `book_request`
  MODIFY `book_request_id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `bus`
--
ALTER TABLE `bus`
  MODIFY `busid` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `bussing`
--
ALTER TABLE `bussing`
  MODIFY `bussingid` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=41;

--
-- AUTO_INCREMENT for table `capstone`
--
ALTER TABLE `capstone`
  MODIFY `capstoneid` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `Class`
--
ALTER TABLE `Class`
  MODIFY `ClassId` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `ClassStd`
--
ALTER TABLE `ClassStd`
  MODIFY `ClassStdId` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `entercompetition`
--
ALTER TABLE `entercompetition`
  MODIFY `entercompid` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `overlooking`
--
ALTER TABLE `overlooking`
  MODIFY `overlookingid` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `parent`
--
ALTER TABLE `parent`
  MODIFY `parentId` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `permission`
--
ALTER TABLE `permission`
  MODIFY `permissionid` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `recordstdcapstone`
--
ALTER TABLE `recordstdcapstone`
  MODIFY `recordstdcapstoneid` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `recordstduni`
--
ALTER TABLE `recordstduni`
  MODIFY `recordstduniid` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `recordstdviolation`
--
ALTER TABLE `recordstdviolation`
  MODIFY `recordstdviolationid` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `recordteaviolation`
--
ALTER TABLE `recordteaviolation`
  MODIFY `recordteaviolationid` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `recordvacation`
--
ALTER TABLE `recordvacation`
  MODIFY `recordvacationid` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `romming`
--
ALTER TABLE `romming`
  MODIFY `rommingid` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `room`
--
ALTER TABLE `room`
  MODIFY `roomid` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `scholarship`
--
ALTER TABLE `scholarship`
  MODIFY `scolarshipid` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `school`
--
ALTER TABLE `school`
  MODIFY `schoolid` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `socialsharing`
--
ALTER TABLE `socialsharing`
  MODIFY `socialsharingid` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `stem`
--
ALTER TABLE `stem`
  MODIFY `stemid` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `student`
--
ALTER TABLE `student`
  MODIFY `StdId` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=191;

--
-- AUTO_INCREMENT for table `student_old`
--
ALTER TABLE `student_old`
  MODIFY `studentid2` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=124;

--
-- AUTO_INCREMENT for table `teachers`
--
ALTER TABLE `teachers`
  MODIFY `teachersId` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=44;

--
-- AUTO_INCREMENT for table `trip`
--
ALTER TABLE `trip`
  MODIFY `tripid` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2520084;

--
-- AUTO_INCREMENT for table `violations`
--
ALTER TABLE `violations`
  MODIFY `violationid` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=23;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `absence`
--
ALTER TABLE `absence`
  ADD CONSTRAINT `absence_ibfk_1` FOREIGN KEY (`stdId`) REFERENCES `student` (`StdId`);

--
-- Constraints for table `AssignScholar`
--
ALTER TABLE `AssignScholar`
  ADD CONSTRAINT `AssignScholar_ibfk_1` FOREIGN KEY (`StdId`) REFERENCES `student` (`StdId`),
  ADD CONSTRAINT `AssignScholar_ibfk_2` FOREIGN KEY (`ScholarshipId`) REFERENCES `scholarship` (`scolarshipid`);

--
-- Constraints for table `bussing`
--
ALTER TABLE `bussing`
  ADD CONSTRAINT `bussing_ibfk_1` FOREIGN KEY (`busid`) REFERENCES `bus` (`busid`),
  ADD CONSTRAINT `bussing_ibfk_2` FOREIGN KEY (`stdid`) REFERENCES `student` (`StdId`);

--
-- Constraints for table `ClassStd`
--
ALTER TABLE `ClassStd`
  ADD CONSTRAINT `ClassStd_ibfk_1` FOREIGN KEY (`ClassId`) REFERENCES `Class` (`ClassId`),
  ADD CONSTRAINT `ClassStd_ibfk_2` FOREIGN KEY (`StdId`) REFERENCES `student` (`StdId`),
  ADD CONSTRAINT `ClassStd_ibfk_3` FOREIGN KEY (`ClassId`) REFERENCES `Class` (`ClassId`),
  ADD CONSTRAINT `ClassStd_ibfk_4` FOREIGN KEY (`StdId`) REFERENCES `student` (`StdId`);

--
-- Constraints for table `entercompetition`
--
ALTER TABLE `entercompetition`
  ADD CONSTRAINT `entercompetition_ibfk_1` FOREIGN KEY (`competitionid`) REFERENCES `competition` (`competitionid`),
  ADD CONSTRAINT `entercompetition_ibfk_2` FOREIGN KEY (`stdid`) REFERENCES `student` (`StdId`);

--
-- Constraints for table `parent`
--
ALTER TABLE `parent`
  ADD CONSTRAINT `parent_ibfk_1` FOREIGN KEY (`studentId`) REFERENCES `student` (`StdId`);

--
-- Constraints for table `recordstdcapstone`
--
ALTER TABLE `recordstdcapstone`
  ADD CONSTRAINT `recordstdcapstone_ibfk_1` FOREIGN KEY (`capstoneid`) REFERENCES `capstone` (`capstoneid`),
  ADD CONSTRAINT `recordstdcapstone_ibfk_2` FOREIGN KEY (`studentid`) REFERENCES `student` (`StdId`),
  ADD CONSTRAINT `recordstdcapstone_ibfk_3` FOREIGN KEY (`studentid2`) REFERENCES `student` (`StdId`),
  ADD CONSTRAINT `recordstdcapstone_ibfk_4` FOREIGN KEY (`studentid3`) REFERENCES `student` (`StdId`),
  ADD CONSTRAINT `recordstdcapstone_ibfk_5` FOREIGN KEY (`studentid4`) REFERENCES `student` (`StdId`),
  ADD CONSTRAINT `recordstdcapstone_ibfk_6` FOREIGN KEY (`studentid5`) REFERENCES `student` (`StdId`),
  ADD CONSTRAINT `recordstdcapstone_ibfk_7` FOREIGN KEY (`studentid6`) REFERENCES `student` (`StdId`);

--
-- Constraints for table `recordstduni`
--
ALTER TABLE `recordstduni`
  ADD CONSTRAINT `recordstduni_ibfk_1` FOREIGN KEY (`stdid`) REFERENCES `student` (`StdId`),
  ADD CONSTRAINT `recordstduni_ibfk_2` FOREIGN KEY (`uniid`) REFERENCES `University` (`UniversityID`);

--
-- Constraints for table `recordstdviolation`
--
ALTER TABLE `recordstdviolation`
  ADD CONSTRAINT `recordstdviolation_ibfk_1` FOREIGN KEY (`studentid`) REFERENCES `student` (`StdId`),
  ADD CONSTRAINT `recordstdviolation_ibfk_2` FOREIGN KEY (`violationid`) REFERENCES `violations` (`violationid`);

--
-- Constraints for table `recordteaviolation`
--
ALTER TABLE `recordteaviolation`
  ADD CONSTRAINT `recordteaviolation_ibfk_2` FOREIGN KEY (`teaviolationname`) REFERENCES `violations` (`violationid`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  ADD CONSTRAINT `recordteaviolation_ibfk_3` FOREIGN KEY (`teacherid`) REFERENCES `teachers` (`teachersId`) ON DELETE RESTRICT ON UPDATE RESTRICT;

--
-- Constraints for table `recordvacation`
--
ALTER TABLE `recordvacation`
  ADD CONSTRAINT `recordvacation_ibfk_2` FOREIGN KEY (`vacationtype`) REFERENCES `addvacation` (`addvacationid`),
  ADD CONSTRAINT `recordvacation_ibfk_3` FOREIGN KEY (`teacherid`) REFERENCES `teachers` (`teachersId`) ON DELETE RESTRICT ON UPDATE RESTRICT;

--
-- Constraints for table `romming`
--
ALTER TABLE `romming`
  ADD CONSTRAINT `romming_ibfk_1` FOREIGN KEY (`roomid`) REFERENCES `room` (`roomid`),
  ADD CONSTRAINT `romming_ibfk_2` FOREIGN KEY (`stdid`) REFERENCES `student` (`StdId`);

--
-- Constraints for table `socialsharing`
--
ALTER TABLE `socialsharing`
  ADD CONSTRAINT `socialsharing_ibfk_1` FOREIGN KEY (`parentid`) REFERENCES `student` (`StdId`);

--
-- Constraints for table `stem`
--
ALTER TABLE `stem`
  ADD CONSTRAINT `stem_ibfk_1` FOREIGN KEY (`groupid`) REFERENCES `capstone` (`capstoneid`),
  ADD CONSTRAINT `stem_ibfk_2` FOREIGN KEY (`stdclass`) REFERENCES `class` (`classid`),
  ADD CONSTRAINT `stem_ibfk_3` FOREIGN KEY (`studentid`) REFERENCES `student` (`StdId`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
