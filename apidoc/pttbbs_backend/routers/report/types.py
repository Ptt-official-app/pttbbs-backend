

from ..comment.types import CommentReportView
from ..community.types import CommunityReportView
from ..message.types import PrivateMessageReportView
from ..post.types import PostReportView


class PostCombinedReportView(PostReportView):
    type_: str = 'post'


class CommentCombinedReportView(CommentReportView):
    type_: str = 'comment'


class PrivateMessageCombinedReportView(PrivateMessageReportView):
    type_: str = 'private_message'


class CommunityCombinedReportView(CommunityReportView):
    type_: str = 'community'


type ReportCombinedView = PostCombinedReportView | CommentCombinedReportView | PrivateMessageCombinedReportView | CommunityCombinedReportView  # noqa
