#!/usr/bin/env sh

docker-compose -f docker-compose-integration-tests.yml up -d db paymentsvc
docker-compose -f docker-compose-integration-tests.yml up --exit-code-from test test
test_result=$?
docker-compose -f docker-compose-integration-tests.yml logs paymentsvc
docker-compose -f docker-compose-integration-tests.yml down --rmi local
exit $test_result
