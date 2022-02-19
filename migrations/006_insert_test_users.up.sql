-- Test User (Identified)
INSERT INTO users(
    id,
    username,
    full_name, 
    email,
    password,
    access_token,
    refresh_token,
    is_identified) VALUES(
    '07040879-12ed-4a75-87d4-269d4f693f46',
    'toshkentov',
    'sardor toshkentov',
    'sardortoshkentov@gmail.com',
    'sardor@11T',
    'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzIjoxNjQ5NTkwMzQzLCJpZCI6IjA3MDQwODc5LTEyZWQtNGE3NS04N2Q0LTI2OWQ0ZjY5M2Y0NiIsInJvbGUiOiJ1c2VyIn0.C54HWjOL2xcfpnLJf_Dy9vuAFovPOvarfz4VW_30Jvo',
    '9028cca878ca251015d8c44f67709667ca90ca919dc8e1ef7147a09991da017f.1646134343',
    true);



-- Test User (Unidentified)

INSERT INTO users(id, username, password, access_token, refresh_token)
    VALUES(
        '2489caa4-bbfe-4015-a0a3-3ff867d5a701',
        'sardor',
        'toshkentov@11T',
        'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzIjoxNjQ5NTkwMzgyLCJpZCI6IjI0ODljYWE0LWJiZmUtNDAxNS1hMGEzLTNmZjg2N2Q1YTcwMSIsInJvbGUiOiJ1c2VyIn0.qWRBeGFCenvNtAAG4IXI9FpVg_SnjAo6d0L3r0Ymnok',
        'e5e1a02a282ae20c95453f5469209d39ef5ee4a8164b2d07401da2de3b93c077.1646134382');