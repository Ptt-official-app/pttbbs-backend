# -*- coding: utf-8 -*-

import yaml
from flask import Flask

app = Flask(__name__.split('.')[0])

_APP_PREFIX = '/api'


def _with_app_prefix(path):
    """Summary

    Args:
        path (TYPE): Description

    Returns:
        TYPE: Description
    """
    if not _APP_PREFIX:
        return path
    return _APP_PREFIX + path


with open('apidoc/template.json', 'r') as f:
    template = yaml.full_load(f)


@app.route('/')
def _index():
    """
    swagger_from_file: apidoc/index.yaml

    """
    return ''


@app.route(_with_app_prefix('/boards/popular'))
def _load_popular_boards():
    """
    swagger_from_file: apidoc/load_popular_boards.yaml

    """
    return ''


@app.route(_with_app_prefix('/boards/refresh-white-list'))
def _refresh_white_list():
    """
    swagger_from_file: apidoc/refresh_white_list.yaml

    """
    return ''


@app.route(_with_app_prefix('/userid'))
def _get_user_id():
    """
    swagger_from_file: apidoc/get_user_id.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/articles'))
def _load_general_articles(bid):
    """
    swagger_from_file: apidoc/load_general_articles.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/articles/bottom'))
def _load_bottom_articles(bid):
    """
    swagger_from_file: apidoc/load_bottom_articles.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/article/<aid>'))
def _get_article_detail(bid, aid):
    """
    swagger_from_file: apidoc/get_article_detail.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/articleblocks/<aid>'))
def _get_article_blocks(bid, aid):
    """
    swagger_from_file: apidoc/get_article_blocks.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/article/<aid>/comments'))
def _load_article_comments(bid, aid):
    """
    swagger_from_file: apidoc/load_article_comments.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>'))
def _get_board_detail(bid):
    """
    swagger_from_file: apidoc/get_board_detail.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/summary'))
def _get_board_summary(bid):
    """
    swagger_from_file: apidoc/get_board_summary.yaml

    """
    return ''


@app.route(_with_app_prefix('/version'), methods=['GET'])
def _get_verion():
    """
    swagger_from_file: apidoc/get_version.yaml

    """
    return ''
