#!/usr/bin/env python

import argparse
import json

from .main import app


def generate_openapi(out_filename: str, host: str):
    openapi_schema = app.openapi()

    openapi_schema["servers"] = [
        {
            'url': f'{host}'
        }
    ]

    if 'components' not in openapi_schema:
        openapi_schema['components'] = {}
    if 'securitySchemes' not in openapi_schema['components']:
        openapi_schema['components']['securitySchemes'] = {
            'default': {
                'type': 'http',
                'scheme': 'bearer',
            },
        }
    if 'security' not in openapi_schema:
        openapi_schema['security'] = [
            {
                'default': [],
            },
        ]

    with open(out_filename, "w") as f:
        json.dump(openapi_schema, f, indent=2)


def _parse_arg():
    parser = argparse.ArgumentParser()
    parser.add_argument('--out')
    parser.add_argument('--host')
    args = parser.parse_args()

    return args


def _main():
    args = _parse_arg()
    generate_openapi(args.out, args.host)


if __name__ == "__main__":
    _main()
