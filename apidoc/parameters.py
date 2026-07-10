#!/usr/bin/env python
# -*- coding: utf-8 -*-

import yaml
import sys
import re
import json
import logging

filename2 = sys.argv[1]
filename3 = sys.argv[2]

with open(filename2, 'r') as f:
    the_struct2 = json.load(f)

with open(filename3, 'r') as f:
    the_struct3 = json.load(f)

paths2 = the_struct2.get('paths', {})
paths3 = the_struct3.get('paths', {})


def _is_not_authorization(x):
    if '$ref' in x:
        return x['$ref'] != '#/definitions/ParamsAuthorization'

    return x['name'] != 'Authorization'


# params
for path, data_by_path in paths3.items():
    for method, data_by_method in data_by_path.items():
        params2 = paths2.get(path, {}).get(method, {}).get('parameters')
        if params2 is None:
            logging.warning('no params: path: %s', path)
            continue
        params2 = list(filter(_is_not_authorization, params2))
        for each in params2:
            if '$ref' in each:
                val = re.sub(r'#/definitions', '#/components/schemas', each['$ref'])
                each['$ref'] = val
            if 'schema' in each:
                schema = each['schema']
                if '$ref' in schema:
                    val = re.sub(r'#/definitions', '#/components/schemas', schema['$ref'])
                    schema['$ref'] = val
        data_by_method['parameters'] = params2

# definitions
definitions2 = the_struct2.get('definitions', {})
schemas3 = the_struct3.get('components', {}).get('schemas', {})
for key, val in definitions2.items():
    if 'name' in val:
        schemas3[key] = val

# servers
the_host = the_struct2.get('host', '')
the_host3 = ''
if the_host.startswith('localhost'):
    the_host3 = 'http://' + the_host
else:
    the_host3 = 'https://' + the_host
the_struct3['servers'] = [{'url': the_host3}]

out_filename3 = re.sub(r'\.json', '.params.json', filename3)
with open(out_filename3, 'w') as f:
    json.dump(the_struct3, f, indent=4)
