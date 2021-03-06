# Druid Manager

![test-go-code](https://github.com/humit0/druid_manager/actions/workflows/test-go-code.yml/badge.svg)

## 소개

Druid에 대한 내용을 모니터링 및 여러 내용을 관리하기 위한 서버입니다.

## 주요 기능

### 데이터 소스 모니터링

1. 데이터 소스 목록
2. 각 데이터 소스 별 총 열 수, 세그먼트 수, 용량
3. 각 데이터 소스 별 차원 및 측정 값
4. 각 데이터 소스 별 롤업 비율 (count 컬럼이 필요함)

### 데이터 소스 알람 규칙

1. 데이터 개수 비교 (count 컬럼이 필요함)

### 쿼리 모니터링

1. 쿼리 타입(sql, native) 별 총 요청 수
2. 성공 수와 실패 수
3. 데이터 소스 별 호출 횟수

### 유저 관리

1. Role 추가 및 목록 확인
2. Role 권한 확인 및 수정
3. 그룹 추가 및 목록 확인
4. 그룹에 Role 및 유저 추가/삭제
5. 유저 목록 및 추가
6. 유저 비밀번호 업데이트 기능

## 문서

- [DB 설계](docs/database_design.md)

## TODO

[todo](todo.md)