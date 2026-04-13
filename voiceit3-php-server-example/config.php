<?php
// Example server configuration. All values must be supplied via env vars
// — there are no production-looking defaults so misconfiguration fails
// loudly instead of silently shipping demo credentials.

function require_env(string $name): string {
    $v = getenv($name, true);
    if ($v === false || $v === '') {
        throw new RuntimeException(
            "Required environment variable {$name} is not set. " .
            "See .env.example for the full list."
        );
    }
    return $v;
}

$VOICEIT_API_KEY     = require_env('VOICEIT_THREE_API_KEY');
$VOICEIT_API_TOKEN   = require_env('VOICEIT_THREE_API_TOKEN');
$VOICEIT_TEST_USERID = require_env('VOICEIT_TEST_USER_ID');
$DEMO_EMAIL          = require_env('DEMO_EMAIL');
$DEMO_PASSWORD       = require_env('DEMO_PASSWORD');
$CONTENT_LANGUAGE    = getenv('CONTENT_LANGUAGE', true) ?: 'no-STT';
