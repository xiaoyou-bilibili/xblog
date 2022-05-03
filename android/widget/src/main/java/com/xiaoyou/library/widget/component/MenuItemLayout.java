package com.xiaoyou.library.widget.component;

import android.content.Context;
import android.content.res.TypedArray;
import android.graphics.drawable.Drawable;
import android.util.AttributeSet;
import android.view.LayoutInflater;
import android.view.View;
import android.widget.FrameLayout;
import android.widget.ImageView;
import android.widget.LinearLayout;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;

import com.xiaoyou.library.widget.R;

import static com.xiaoyou.library.widget.util.DensityUtilKt.dip2px;


public class MenuItemLayout extends FrameLayout {
    private static final String TAG = "MenuItemLayout";
    private Context mContext;
    private View view;
    private TextView main_text, hint_text, xian;
    private ImageView text_img, more, more_right;
    private OnClickListener onClickListener;
    private String titleText;
    private String hintText;
    private int textImgId;
    private String onclickId;
    private LinearLayout linearLayout;
    public static final int NO_LINE = 0;
    public static final int DIVIDE_LINE = 1;
    public static final int DIVIDE_AREA = 2;
    public int divideLineStyle = NO_LINE;
    private boolean isShowRedHintImg = false;

    public MenuItemLayout(@NonNull Context context) {
        this(context, null);
    }

    public MenuItemLayout(@NonNull Context context, @Nullable AttributeSet attrs) {
        this(context, attrs, 0);
    }

    public MenuItemLayout(@NonNull Context context, @Nullable AttributeSet attrs, int defStyleAttr) {
        super(context, attrs, defStyleAttr);
        init(context, attrs);
    }

    private void init(Context context, AttributeSet attrs) {
        mContext = context;
        LayoutInflater inflater = (LayoutInflater) mContext.getSystemService(Context.LAYOUT_INFLATER_SERVICE);
        view = inflater.inflate(R.layout.item_menu_layout, this, true);
        main_text = view.findViewById(R.id.text);
        hint_text = view.findViewById(R.id.text_hint);
        text_img = view.findViewById(R.id.text_img);
        more = view.findViewById(R.id.more);
        more_right = view.findViewById(R.id.more_right);
        linearLayout = view.findViewById(R.id.main_lin);
        xian = view.findViewById(R.id.xian);

        TypedArray a = mContext.obtainStyledAttributes(attrs, R.styleable.MenuItemLayout);
        setTitleText(a.getString(R.styleable.MenuItemLayout_TitleText));
        setHint_text(a.getString(R.styleable.MenuItemLayout_PromptText));
        setIconImgId(a.getResourceId(R.styleable.MenuItemLayout_TitleImg, 10000));
        isSwitchmore(a.getBoolean(R.styleable.MenuItemLayout_isSwitch, false));
        setLinBackground(a.getResourceId(R.styleable.MenuItemLayout_Background, -1));
        setImageLeftMargin(a.getInt(R.styleable.MenuItemLayout_ImageLeftMargin, 0));
        isUnderline(a.getBoolean(R.styleable.MenuItemLayout_Underline, true));
        setTitleSize(a.getInt(R.styleable.MenuItemLayout_TitleSize, 19));
        setImageSize(a.getInt(R.styleable.MenuItemLayout_ImageSize, 24));
        a.recycle();
    }

    public void setTitleSize(int value) {
        main_text.setTextSize(value);
    }

    public void setImageSize(int value) {
        int px = dip2px(mContext, value);
        //将用户输入的数据转换为dp
        LinearLayout.LayoutParams linearParams = (LinearLayout.LayoutParams) text_img.getLayoutParams();
        //取控件imageView当前的布局参数 linearParams.height/width = value;// 控件的高强制设成用户设置的
        linearParams.width = px;
        linearParams.height = px;
        // 控件的宽强制设成30
        text_img.setLayoutParams(linearParams);
        //使设置好的布局参数应用到控件
    }

    public void isUnderline(boolean value) {
        if (!value) {
            xian.setVisibility(GONE);
        } else {
            xian.setVisibility(VISIBLE);
        }
    }

    public void setImageLeftMargin(int value) {
        setMargins(text_img, value, 0, 0, 0);
    }

    public int getIconImgId() {
        return textImgId;
    }

    public void setIconImgId(int iconImgId) {
        if (iconImgId != 10000) {
            this.textImgId = iconImgId;
            text_img.setImageResource(iconImgId);
        } else {
            text_img.setVisibility(View.GONE);
        }
    }

    public String getTitleText() {
        return titleText;
    }

    public void setTitleText(String titleText) {
        if (titleText != null) {
            this.titleText = titleText;
            main_text.setText(titleText);
        }
    }

    public String getHintText() {
        return hintText;
    }

    public void setHint_text(String text) {
        if (text != null) {
            this.hintText = text;
            hint_text.setText(text);
        } else {
            hint_text.setVisibility(GONE);
        }
    }

    public void isSwitchmore(boolean is) {
        if (is) {
            more.setVisibility(View.VISIBLE);
            more_right.setVisibility(View.GONE);
        } else {
            more.setVisibility(View.GONE);
            more_right.setVisibility(View.VISIBLE);
        }
    }

    @Override
    public void setBackground(Drawable background) {
        linearLayout.setBackground(background);
    }

    public void setLinBackground(int reference) {
        if (reference == -1) {
            return;
        }
        linearLayout.setBackgroundResource(reference);
    }


    public void setViewOnlickListener(OnClickListener onlickListener) {
        this.onClickListener = onlickListener;
        view.setOnClickListener(onlickListener);
    }

    public TextView getTitleTv() {
        return main_text;
    }

    public TextView getHintTv() {
        return hint_text;
    }

    public static void setMargins(View v, int l, int t, int r, int b) {
        if (v.getLayoutParams() instanceof MarginLayoutParams) {
            MarginLayoutParams p = (MarginLayoutParams) v.getLayoutParams();
            p.setMargins(l, t, r, b);
            v.requestLayout();
        }
    }
}
