from typing import Any, Literal

from pydantic import BaseModel, Field

from ..types import ID, DbURL, Time3339
from ..user.types import PersonView


class PluginMeta(BaseModel):
    description: str = ''
    url: str = ''
    name: str


class LocalSiteUrl(BaseModel):
    updated_at: Time3339 = ''
    published_at: Time3339
    url: str
    id: ID


class AdminOAuthProvider(BaseModel):
    '''
    admin-oauth-provider (pttbbs-backend as OP)
    '''
    use_pkce: bool
    updated_at: Time3339 = ''
    published_at: Time3339
    enabled: bool
    account_linking_enabled: bool
    auto_verify_email: bool
    scopes: str = Field(description='Lists the scopes requested from users. Users will have to grant access to the requested scope at sign up.')  # noqa
    client_id: str
    id_claim: str = Field(description='The OAuth 2.0 claim containing the unique user ID returned by the provider. Usually this should be set to "sub".')  # noqa
    userinfo_endpoint: str = Field(description='The UserInfo Endpoint is an OAuth 2.0 Protected Resource that returns Claims about the authenticated End-User. This is defined in the OIDC specification.')  # noqa
    token_endpoint: str = Field(description='The token endpoint is used by the client to obtain an access token by presenting its authorization grant or refresh token. This is usually provided by the OAUTH provider.')  # noqa
    authorization_endpoint: str = Field(description='The authorization endpoint is used to interact with the resource owner and obtain an authorization grant. This is usually provided by the OAUTH provider.')  # noqa
    issuer: str = Field(description='The issuer url of the OAUTH provider.')
    display_name: str = Field(description='The OAuth 2.0 provider name displayed to the user on the Login page.')  # noqa
    id: ID


class PublicOAuthProvider(BaseModel):
    '''
    oauth provider (ptbbs-backend as RP)
    '''
    use_pkce: bool
    scopes: str = Field(description='Lists the scopes requested from users. Users will have to grant access to the requested scope at sign up.')  # noqa
    client_id: str
    authorization_endpoint: str = Field(description='The authorization endpoint is used to interact with the resource owner and obtain an authorization grant. This is usually provided by the OAUTH provider.')  # noqa
    display_name: str
    id: ID


class TagLine(BaseModel):
    updated_at: Time3339 = ''
    published_at: Time3339
    content: str
    id: ID


class Language(BaseModel):
    name: str
    code: str
    id: ID


class Instance(BaseModel):
    version: str
    software: str
    updated_at: Time3339 = ''
    published_at: Time3339
    domain: str
    id: ID


class LocalSiteRateLimit(BaseModel):
    import_user_settings_interval_seconds: float
    import_user_settings_max_requests: int
    updated_at: Time3339 = ''
    published_at: Time3339
    search_interval_seconds: float
    search_max_requests: int
    comment_interval_seconds: float
    comment_max_requests: int
    image_interval_seconds: float
    image_max_requests: int
    register_interval_seconds: float
    register_max_requests: int
    post_interval_seconds: float
    post_max_requests: int
    message_interval_seconds: float
    message_max_requests: int
    local_site_id: ID


class CreateLocalSiteRateLimit(BaseModel):
    rate_limit_import_user_settings_interval_seconds: float = 0
    rate_limit_import_user_settings_max_requests: int = 0
    rate_limit_search_interval_seconds: float = 0
    rate_limit_search_max_requests: int = 0
    rate_limit_comment_interval_seconds: float = 0
    rate_limit_comment_max_requests: int = 0
    rate_limit_image_interval_seconds: float = 0
    rate_limit_image_max_requests: int = 0
    rate_limit_register_interval_seconds: float = 0
    rate_limit_register_max_requests: int = 0
    rate_limit_post_interval_seconds: float = 0
    rate_limit_post_max_requests: int = 0
    rate_limit_message_interval_seconds: float = 0
    rate_limit_message_max_requests: int = 0


type ImageMode = Literal['none', 'store_link_previews', 'proxy_all_images']

type FederationMode = Literal['all', 'local', 'disable']

type CommentSortType = Literal['hot', 'top', 'new', 'old', 'controversial']

type PostSortType = Literal[
    'active',
    'hot',
    'new',
    'old',
    'top',
    'most_comments',
    'new_comments',
    'controversial',
    'scaled']

type PostListingMode = Literal['list', 'card', 'small_card']

type RegistrationMode = Literal[
    'closed', 'require_application', 'require_invitation', 'open']

type ListingType = Literal[
    'all', 'local', 'subscribed', 'moderator_view', 'suggested']


class LocalSiteConfig(BaseModel):
    max_invites_per_user_allowed: int = Field(description='How many active invite links a user can have')  # noqa
    image_upload_disabled: bool
    image_allow_video_uploads: bool = Field(description='This affects post and comment images, but not avatars and banners.')  # noqa
    image_max_upload_size: int = Field(description='This affects post and comment images, but not avatar and banner sizes.')  # noqa
    image_max_banner_size: int
    image_max_avatar_size: int
    image_max_thumbnail_size: int = Field(description='These are pixel sizes. Larger images are automatically downscaled.')  # noqa
    image_upload_timeout_seconds: float
    image_proxy_bypass_domains: str = Field(default='', description='Allows bypassing proxy for specific image hosts when using [[ImageMode.ProxyAllImages]]. Use a comma-delimited string.')  # noqa
    image_mode: ImageMode = Field(
        description='A mode for setting how pictrs handles images.')

    suggested_multi_community_id: ID
    email_notifications_disabled: bool
    nsfw_content_disallowed: bool = Field(
        description='Block NSFW content being created')
    comment_downvotes: FederationMode
    comment_upvotes: FederationMode
    post_downvotes: FederationMode
    post_upvotes: FederationMode

    federation_signed_fetch: bool = Field(description='Whether to sign outgoing Activitypub fetches with private key of local instance. Some Fediverse instances and platforms require this.')  # noqa
    reports_email_admins: bool = Field(
        description='Whether to email admins on new reports.')

    oauth_registration: bool = Field(description='Whether or not external auth methods can auto-register users.')  # noqa
    registration_mode: RegistrationMode = Field(description='The registration mode for your site. Determines what happens after a user signs up.')  # noqa

    federation_enabled: bool

    updated_at: Time3339 = ''
    published_at: Time3339

    slur_filter_regex: str = Field(
        default='', description='An optional regex to filter words.')

    application_email_admins: bool = Field(
        description='Whether new applications email admins.')

    legal_information: str = ''

    default_comment_sort_type: CommentSortType
    default_items_per_page: int
    default_post_time_range_seconds: int = Field(
        default=0, description='A default time range limit to apply to post sorts, in seconds.')
    default_post_sort_type: PostSortType
    default_post_listing_mode: PostListingMode = Field(description='A post-view mode that changes how multiple post listings look.')  # noqa
    default_post_listing_type: ListingType = Field(description='A listing type for post and comment list fetches.')  # noqa
    default_theme: str

    private_instance: bool = Field(
        description='Whether the instance is private or public.')
    application_question: str = Field(
        default='', description='An optional registration application questionnaire in markdown.')
    email_verification_required: bool = Field(
        description='Whether emails are required.')
    community_creation_admin_only: bool = Field(description='Whether only admins can create communities.')  # noqa


class PartialLocalSiteConfig(BaseModel):
    max_invites_per_user_allowed: int = Field(default=0, description='How many active invite links a user can have')  # noqa
    image_upload_disabled: bool = False
    image_allow_video_uploads: bool = Field(default=False, description='This affects post and comment images, but not avatars and banners.')  # noqa
    image_max_upload_size: int = Field(default=0, description='This affects post and comment images, but not avatar and banner sizes.')  # noqa
    image_max_banner_size: int = 0
    image_max_avatar_size: int = 0
    image_max_thumbnail_size: int = Field(default=0, description='These are pixel sizes. Larger images are automatically downscaled.')  # noqa
    image_upload_timeout_seconds: float = 0
    image_proxy_bypass_domains: str = Field(default='', description='Allows bypassing proxy for specific image hosts when using [[ImageMode.ProxyAllImages]]. Use a comma-delimited string.')  # noqa
    image_mode: ImageMode = Field(
        default='none', description='A mode for setting how pictrs handles images.')

    suggested_multi_community_id: ID = 0
    email_notifications_disabled: bool = False
    nsfw_content_disallowed: bool = Field(
        default=False, description='Block NSFW content being created')
    comment_downvotes: FederationMode = 'local'
    comment_upvotes: FederationMode = 'local'
    post_downvotes: FederationMode = 'local'
    post_upvotes: FederationMode = 'local'

    federation_signed_fetch: bool = Field(default=False, description='Whether to sign outgoing Activitypub fetches with private key of local instance. Some Fediverse instances and platforms require this.')  # noqa
    reports_email_admins: bool = Field(
        default=False, description='Whether to email admins on new reports.')

    oauth_registration: bool = Field(default=False, description='Whether or not external auth methods can auto-register users.')  # noqa
    registration_mode: RegistrationMode = Field(default='open', description='The registration mode for your site. Determines what happens after a user signs up.')  # noqa

    federation_enabled: bool = False

    slur_filter_regex: str = Field(
        default='', description='An optional regex to filter words.')

    application_email_admins: bool = Field(
        default=False, description='Whether new applications email admins.')

    legal_information: str = ''

    default_comment_sort_type: CommentSortType = 'new'
    default_items_per_page: int = 0
    default_post_time_range_seconds: int = Field(
        default=0, description='A default time range limit to apply to post sorts, in seconds.')
    default_post_sort_type: PostSortType = 'new'
    default_post_listing_mode: PostListingMode = Field(default='list', description='A post-view mode that changes how multiple post listings look.')  # noqa
    default_post_listing_type: ListingType = Field(default='local', description='A listing type for post and comment list fetches.')  # noqa
    default_theme: str = ''

    private_instance: bool = Field(
        default=False, description='Whether the instance is private or public.')
    application_question: str = Field(
        default='', description='An optional registration application questionnaire in markdown.')
    email_verification_required: bool = Field(
        default=False, description='Whether emails are required.')
    community_creation_admin_only: bool = Field(default=False, description='Whether only admins can create communities.')  # noqa


class LocalSite(LocalSiteConfig):

    users_active_half_year: int = Field(description='The number of users with any activity in the last half year.')  # noqa
    users_active_month: int = Field(description='The number of users with any activity in the last month.')  # noqa
    users_active_week: int = Field(description='The number of users with any activity in the last week.')  # noqa
    users_active_day: int = Field(description='The number of users with any activity in the last day.')  # noqa
    communities: int
    comments: int
    posts: int
    users: int

    site_setup: bool = Field(description='True if the site is set up.')
    site_id: ID
    id: ID


class SiteConfig(BaseModel):
    content_warning: str = Field(default='', description='If present, nsfw content is visible by default. Should be displayed by frontends/clients when the site is first opened by a user.')  # noqa
    summary: str = Field(
        default='', description='A shorter, one-line summary of the site.')
    sidebar: str = Field(
        default='', description='A sidebar for the site in markdown.')
    name: str


class PartialSiteConfig(SiteConfig):
    name: str = ''


class Site(SiteConfig):
    instance_id: ID
    inbox_url: DbURL
    last_refreshed_at: Time3339
    ap_id: DbURL
    banner: DbURL
    icon: DbURL

    updated_at: Time3339 = ''
    published_at: Time3339

    id: ID


class SiteView(BaseModel):
    instance: Instance
    local_site_rate_limit: LocalSiteRateLimit
    local_site: LocalSite
    site: Site


class SiteInfo(BaseModel):
    captcha_enabled: bool

    last_application_duration_seconds: float = Field(default=0, description='The number of seconds between the last application published, and approved / denied time. Useful for estimating when your application will be approved.')  # noqa

    active_plugins: list[PluginMeta]

    blocked_urls: list[LocalSiteUrl]

    admin_oauth_providers: list[AdminOAuthProvider]
    oauth_providers: list[PublicOAuthProvider]

    tagline: TagLine

    discussion_languages: list[ID]

    all_languages: list[Language]

    version: str

    admins: list[PersonView]

    site_view: SiteView


class CreateSiteParams(SiteConfig, PartialLocalSiteConfig, CreateLocalSiteRateLimit):
    pass


class UpdateSiteParams(PartialSiteConfig, PartialLocalSiteConfig, CreateLocalSiteRateLimit):
    pass


class UploadSiteIconResult(BaseModel):
    filename: str
    image_url: str


class DeleteSiteIconResult(BaseModel):
    success: bool


class UploadSiteBannerResult(BaseModel):
    filename: str
    image_url: str


class DeleteSiteBannerResult(BaseModel):
    success: bool


class NodeInfoServices(BaseModel):
    outbound: list[str] = []
    inbound: list[str] = []


class NodeInfoUsers(BaseModel):
    activeMonth: int = 0
    activeHalfyear: int = 0
    total: int = 0


class NodeInfoUsage(BaseModel):
    localComments: int = 0
    localPosts: int = 0
    users: NodeInfoUsers


class NodeInfoSoftware(BaseModel):
    homepage: str = ''
    repository: str = ''
    version: str = ''
    name: str = ''


class NodeInfo(BaseModel):
    metadata: dict[str, Any] = {}
    services: NodeInfoServices
    openRegistrations: bool = False
    usage: NodeInfoUsage
    protocols: list[str] = []
    software: NodeInfoSoftware
    version: str
