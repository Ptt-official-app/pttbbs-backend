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


@app.route(_with_app_prefix('/account/login'), methods=['POST'])
def _login():
    """
    swagger_from_file: apidoc/login.yaml

    """
    return ''


@app.route(_with_app_prefix('/account/refresh'), methods=['POST'])
def _refresh():
    """
    swagger_from_file: apidoc/refresh.yaml

    """
    return ''


@app.route(_with_app_prefix('/account/logout'), methods=['POST'])
def _logout():
    """
    swagger_from_file: apidoc/logout.yaml

    """
    return ''


@app.route(_with_app_prefix('/account/register'), methods=['POST'])
def _register_user():
    """
    swagger_from_file: apidoc/register_user.yaml

    """
    return ''


@app.route(_with_app_prefix('/account/attemptregister'), methods=['POST'])
def _attempt_register_user():
    """
    swagger_from_file: apidoc/attempt_register_user.yaml

    """
    return ''


@app.route(_with_app_prefix('/account/existsuser'), methods=['POST'])
def _existsuser():
    """
    swagger_from_file: apidoc/exists_user.yaml

    """
    return ''


@app.route(_with_app_prefix('/articles/popular'))
def _load_popular_articles():
    """
    swagger_from_file: apidoc/load_popular_articles.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<username>/favorites'))
def _load_favorite_boards(username):
    """
    swagger_from_file: apidoc/load_favorite_boards.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<username>/favorites/addboard'), methods=['POST'])
def _add_favorite_board(username):
    """
    swagger_from_file: apidoc/add_favorite_board.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<username>/favorites/addfolder'), methods=['POST'])
def _add_favorite_folder(username):
    """
    swagger_from_file: apidoc/add_favorite_folder.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<username>/favorites/addline'), methods=['POST'])
def _add_favorite_line(username):
    """
    swagger_from_file: apidoc/add_favorite_line.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<username>/favorites/delete'), methods=['POST'])
def _delete_favorite(username):
    """
    swagger_from_file: apidoc/delete_favorite.yaml

    """
    return ''


@app.route(_with_app_prefix('/boards'))
def _load_general_boards():
    """
    swagger_from_file: apidoc/load_general_boards.yaml
    """
    return ''


@app.route(_with_app_prefix('/boards/byclass'))
def _load_general_boards_by_class():
    """
    swagger_from_file: apidoc/load_general_boards_by_class.yaml
    """
    return ''


@app.route(_with_app_prefix('/boards/autocomplete'))
def _load_auto_complete_boards():
    """
    swagger_from_file: apidoc/load_auto_complete_boards.yaml
    """
    return ''


@app.route(_with_app_prefix('/boards/popular'))
def _load_popular_boards():
    """
    swagger_from_file: apidoc/load_popular_boards.yaml

    """
    return ''


@app.route(_with_app_prefix('/cls/<clsid>'))
def _load_class_boards():
    """
    swagger_from_file: apidoc/load_class_boards.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<username>'))
def _get_user_info(username):
    """
    swagger_from_file: apidoc/get_user_info.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<username>/summary'))
def _get_user_summary(username):
    """
    swagger_from_file: apidoc/get_user_summary.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<username>/updatepasswd'), methods=['POST'])
def _change_passwd(username):
    """
    swagger_from_file: apidoc/change_passwd.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<username>/articles'))
def _load_user_articles(username):
    """
    swagger_from_file: apidoc/load_user_articles.yaml

    Args:
        username (TYPE): Description

    Returns:
        TYPE: Description
    """
    return ''


@app.route(_with_app_prefix('/user/<username>/comments'))
def _load_user_comments(username):
    """
    swagger_from_file: apidoc/load_user_comments.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<username>/attemptchangeemail'), methods=['POST'])
def _attempt_change_email(username):
    """
    swagger_from_file: apidoc/attempt_change_email.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<username>/changeemail'), methods=['POST'])
def _change_email(username):
    """
    swagger_from_file: apidoc/change_email.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<username>/attemptsetidemail'), methods=['POST'])
def _attempt_set_id_email(username):
    """
    swagger_from_file: apidoc/attempt_set_id_email.yaml

    """
    return ''


@app.route(_with_app_prefix('/user/<username>/setidemail'), methods=['POST'])
def _set_id_email(username):
    """
    swagger_from_file: apidoc/set_id_email.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/manuals'))
def _load_man_articles(bid):
    """
    swagger_from_file: apidoc/load_man_articles.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/manual/<path:aid>'))
def _get_man_article_detail(bid, aid):
    """
    swagger_from_file: apidoc/get_man_article_detail.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/manualblocks/<path:aid>'))
def _get_man_article_blocks(bid, aid):
    """
    swagger_from_file: apidoc/get_man_article_blocks.yaml

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


@app.route(_with_app_prefix('/board/<bid>/article'), methods=['POST'])
def _create_article(bid):
    """
    swagger_from_file: apidoc/create_article.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/deletearticles'), methods=['POST'])
def _delete_articles(bid):
    """
    swagger_from_file: apidoc/delete_articles.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/article/<aid>/comments'))
def _load_article_comments(bid, aid):
    """
    swagger_from_file: apidoc/load_article_comments.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/article/<aid>/comment'), methods=['POST'])
def _create_comment(bid, aid):
    """
    swagger_from_file: apidoc/create_comment.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/article/<aid>/crosspost'), methods=['POST'])
def _cross_post(bid, aid):
    """
    swagger_from_file: apidoc/cross_post.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/article/<aid>/rank'), methods=['POST'])
def _create_rank(bid, aid):
    """
    swagger_from_file: apidoc/create_rank.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/article/<aid>/edit'), methods=['POST'])
def _edit_article(bid, aid):
    """
    swagger_from_file: apidoc/edit_article.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/article/<aid>/replycomments'), methods=['POST'])
def _replycomments(bid, aid):
    """
    swagger_from_file: apidoc/reply_comments.yaml

    """
    return ''


@app.route(_with_app_prefix('/board/<bid>/article/<aid>/deletecomments'), methods=['POST'])
def _deletecomments(bid, aid):
    """
    swagger_from_file: apidoc/delete_comments.yaml

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


@app.route(_with_app_prefix('/board/<bid>/users'))
def _get_board_users(bid):
    """
    swagger_from_file: apidoc/load_board_users.yaml

    """
    return ''


@app.route(_with_app_prefix('/class/<cls>/board'), methods=['POST'])
def _create_board(bid):
    """
    swagger_from_file: apidoc/create_board.yaml

    """
    return ''


@app.route(_with_app_prefix('/version'), methods=['GET'])
def _get_verion():
    """
    swagger_from_file: apidoc/get_version.yaml

    """
    return ''


@app.route(_with_app_prefix('/uservisitcount'), methods=['GET'])
def _get_user_visit_count():
    """
    swagger_from_file: apidoc/get_user_visit_count.yaml

    """
    return ''
