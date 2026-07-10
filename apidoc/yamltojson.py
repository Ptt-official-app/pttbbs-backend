#!/usr/bin/env python
# -*- coding: utf-8 -*-

import yaml
import json
import sys
import os
import logging
import re

src = sys.argv[1]
dst = sys.argv[2]


def include(self, node):
    filename_list = re.split(r'!include ', self.construct_scalar(node))
    filename_list = [each.strip() for each in filename_list]

    filename_list = [os.sep.join([os.path.dirname(self.name), each]) for each in filename_list]

    logging.warning('include: self: %s filename_list: %s', self.name, filename_list)

    result = {}
    for each in filename_list:
        with open(each, 'r') as f:
            each_struct = yaml.full_load(f)
            result.update(each_struct)
    return result

    '''
    with open(filename, 'r') as f:
        return yaml.full_load(f)
    '''


yaml.add_constructor('!include', include)


with open(src, 'r') as f:
    the_struct = yaml.full_load(f)

with open(dst, 'w') as f:
    json.dump(the_struct, f, indent=4)
